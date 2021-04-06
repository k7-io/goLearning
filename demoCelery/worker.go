package main

import (
	"fmt"
	"github.com/gocelery/gocelery"
	"github.com/gomodule/redigo/redis"
	"reflect"
	"time"
)


func minus(n int) int{
	start := time.Now()
	fmt.Printf("n:%v\n", n)
	decrement(n)
	fmt.Printf("since:%v\n", time.Since(start))
	return n
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
	go func() {
		// celery worker 启动后
		fmt.Printf("go server start after 3s\n")
		time.Sleep(3 * time.Second)
		fmt.Printf("go server ok\n")
		// prepare arguments
		taskName := "main.add"
		// run task
		asyncResult, err := cli.Delay(taskName, 100, 100)
		if err != nil {
			panic(err)
		}
		fmt.Printf("task status:%v\n", asyncResult.TaskID)
		// get results from backend with timeout
		res, err := asyncResult.Get(1 * time.Second)
		if err != nil {
			panic(err)
		}
		fmt.Printf("result: %+v of type %+v", res, reflect.TypeOf(res))
	}()

	// register task
	cli.Register("main.minus", minus)
	cli.Register("main.add", add)

	// start workers (non-blocking call)
	cli.StartWorker()
	fmt.Printf("-->start go work\n")

	// wait for client request
	time.Sleep(1000 * time.Second)

	// stop workers gracefully (blocking call)
	cli.StopWorker()
}