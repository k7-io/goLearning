package main

import (
	"fmt"
	"runtime"
	"time"
)

func Task3()  {
	defer fmt.Println("task3 stop")
	fmt.Println("task3 start")
	fmt.Println("task3 work")
}

func Task4()  {
	defer fmt.Println("task4 stop")
	fmt.Println("task4 start")
	runtime.Goexit() //The effect is the same as return
	fmt.Println("task4 work")
}

func main()  {
	go Task3()
	go Task4()
	time.Sleep(time.Second * 5)
}

