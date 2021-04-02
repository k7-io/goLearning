## py消费者
py实现任务签名，使用python触发任务。但是不需要启动worker.
```py
from celery import Celery

app = Celery('tasks',
    broker='redis://localhost:6379/1',
    backend='redis://localhost:6379/2'
)
app.conf.update(
    CELERY_TASK_SERIALIZER='json',
    CELERY_ACCEPT_CONTENT=['json'],
    CELERY_RESULT_SERIALIZER='json',
    CELERY_ENABLE_UTC=True,
    CELERY_TASK_PROTOCOL=1,
)

@app.on_after_configure.connect
def setup_periodic_tasks(sender, **kwargs):
    """
    需要启动周期任务.
    每5s调度一次minus
    """
    print('-> sender:{}'.format(sender))
    sender.add_periodic_task(5.0, minus.s())

@app.task(name='main.minus')
def minus():
    """
    1亿减到1.
    cpu计算型任务.
    """
    x = 10000000
    # while x > 1:
    #     x -= 1
    print("x ==> done")


@app.task(name='main.add')
def add(x, y):
    return x + y
## celery -A main beat

if __name__ == '__main__':
    print("add:{}".format(add))
    ar = add.apply_async((5456, 2879), serializer='json')
    print("ar:{}".format(ar))
    # print(ar.get())
```
启动周期任务:
```shell
git:(dev_hyh) ✗ celery -A main beat -l debug
-> sender:<Celery tasks at 0x106026990>
celery beat v4.3.0 (rhubarb) is starting.
__    -    ... __   -        _
LocalTime -> 2021-04-02 15:22:35
Configuration ->
    . broker -> redis://localhost:6379/1
    . loader -> celery.loaders.app.AppLoader
    . scheduler -> celery.beat.PersistentScheduler
    . db -> celerybeat-schedule
    . logfile -> [stderr]@%DEBUG
    . maxinterval -> 5.00 minutes (300s)
[2021-04-02 15:22:35,973: DEBUG/MainProcess] Setting default socket timeout to 30
[2021-04-02 15:22:35,973: INFO/MainProcess] beat: Starting...
[2021-04-02 15:22:35,978: WARNING/MainProcess] key:entries
[2021-04-02 15:22:35,978: WARNING/MainProcess] key:tz
[2021-04-02 15:22:35,979: WARNING/MainProcess] key:utc_enabled
[2021-04-02 15:22:35,987: DEBUG/MainProcess] Current schedule:
<ScheduleEntry: main.minus() main.minus() <freq: 5.00 seconds>
[2021-04-02 15:22:35,987: DEBUG/MainProcess] beat: Ticking with max interval->5.00 minutes
[2021-04-02 15:22:35,987: WARNING/MainProcess] key:entries
[2021-04-02 15:22:35,988: DEBUG/MainProcess] beat: Waking up in 4.98 seconds.
[2021-04-02 15:22:40,971: DEBUG/MainProcess] beat: Synchronizing schedule...
[2021-04-02 15:22:40,971: WARNING/MainProcess] key:entries
[2021-04-02 15:22:40,999: INFO/MainProcess] Scheduler: Sending due task main.minus() (main.minus)
[2021-04-02 15:22:41,059: DEBUG/MainProcess] main.minus sent. id->db5ba142-7c80-4c47-a10e-b0a68e5f9a79
[2021-04-02 15:22:41,059: DEBUG/MainProcess] beat: Waking up in 4.91 seconds.
```

启动单个任务. 触发main.add:
```shell
() ➜  demoCelery git:(dev_hyh) ✗ python main.py
-> sender:<Celery tasks at 0x10f71f0d0>
add:<@task: main.add of tasks at 0x10f71f0d0>
ar:0201bd9b-9d70-4fa9-8872-4587d9dd5b37
8335
() ➜  demoCelery git:(dev_hyh) ✗ 
```


## go worker
go worker 注册任务.
* main.minus
* main.add
```go
package main

import (
	"fmt"
	"github.com/gocelery/gocelery"
	"github.com/gomodule/redigo/redis"
	"time"
)


func minus()  {
	start := time.Now()
	fmt.Printf("start:%v\n", start)
	decrement(100000)
	fmt.Printf("end start:%v since:%v\n", start, time.Since(start))
}
func decrement(n int)  {
	if n > 0 {
		n -= 1
	}
}
func add(x, y int) int{
	fmt.Printf("add-->%v + %v\n", x, y)
	return x + y
}


func main() {
	// create redis connection pool
	redisPool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			c, err := redis.DialURL("redis://127.0.0.1/1")
			if err != nil {
				return nil, err
			}
			return c, err
		},
	}

	// initialize celery client
	cli, err := gocelery.NewCeleryClient(
		gocelery.NewRedisBroker(redisPool),
		&gocelery.RedisCeleryBackend{Pool: redisPool},
		5, // number of workers
	)
	if err != nil {
		panic(err)
	}

	// register task
	cli.Register("main.minus", minus)
	cli.Register("main.add", add)

	// start workers (non-blocking call)
	cli.StartWorker()
	fmt.Printf("-->start go work")

	// wait for client request
	time.Sleep(1000 * time.Second)

	// stop workers gracefully (blocking call)
	cli.StopWorker()
}
```
启动worker
```
➜  demoCelery git:(dev_hyh) ✗ go run main.go 
-->start go workstart:2021-04-02 15:26:07.648563 +0800 CST m=+8.071368711
end start:2021-04-02 15:26:07.648563 +0800 CST m=+8.071368711 since:33.972µs
add-->5456 + 2879
```

## go 消费者