package main

import (
	"fmt"
	"time"
)

func main()  {
	/*ch := make(chan int)
	go func() {
		for i := 0; i < 3; i++ {
			ch <- i
		}
	}()
	for {
		select {
		case msg := <-ch:
			fmt.Println(msg)
		default:
			time.Sleep(time.Second)
		}
	}*/

	ch := make(chan int)
	done := make(chan bool)
	go func() {
		for {
			select {
			case msg := <-ch:
				fmt.Println(msg)
			case <-time.After(time.Second * 3):
				fmt.Println("over time")
				done <- true
			}
		}
	}()

	for i := 0; i < 10; i++ {
		ch <- i
	}
	<-done
	fmt.Println("program exit")

	//deadlock example
	//select {
	//
	//}
}