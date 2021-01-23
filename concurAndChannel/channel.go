package main

import (
	"fmt"
	"time"
)

func main()  {
	ch := make(chan string) // make(chan string , 0)
	go func() {
		fmt.Println(<-ch)
	}()
	ch <- "test"

	//unbuffered channel
	ch1 := make(chan int, 0)
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Printf("len(ch1)=%v, cap(ch1)=%v\n", len(ch1), cap(ch1))
			ch1 <- i
		}
	}()
	for i := 0; i < 3; i++ {
		time.Sleep(time.Second)
		fmt.Println(<-ch1)
	}

	//buffered channel
	ch2 := make(chan int, 3)
	go func() {
		for i := 0; i < 3; i++ {
			ch2 <- i
			fmt.Printf("len(ch2)=%v, cap(ch2)=%v\n", len(ch2), cap(ch2))
		}
		close(ch2)
	}()


	//for {
	//	if val, ok := <- ch2; ok == true{
	//		fmt.Println(val)
	//	}else{
	//		return
	//	}
	//}
	for data := range ch2{
		fmt.Println(data)
	}
}