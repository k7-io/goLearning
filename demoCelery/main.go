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