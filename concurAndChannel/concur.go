package main

import (
	"fmt"
	"time"
)

func Task1()  {
	for {
		fmt.Println(time.Now().Format("15:04:05"), "now handle Task1")
		time.Sleep(time.Second * 3)
	}
}

func Task2()  {
	for {
		fmt.Println(time.Now().Format("15:04:05"), "now handle Task2")
		time.Sleep(time.Second * 1)
	}
}

func main()  {
	//go Task1()
	//go Task2()
	//for{
	//	fmt.Println(time.Now().Format("15:04:05"), "now handle Main")
	//	time.Sleep(time.Second * 2)
	//}

	go func() {
		for {
			fmt.Println(time.Now().Format("15:04:05"), "now handle Task1")
			time.Sleep(time.Second * 3)
		}
	}()
	time.Sleep(time.Second * 100)
}
