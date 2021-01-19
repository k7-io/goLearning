package main

import (
	"fmt"
)

func addSub(x, y int) (sum int,sub int) {
	sum = x + y
	sub = x - y
	return sum, sub
}

func main()  {
	a := 1
	b := 2
	sum, sub := addSub(a, b)
	fmt.Println(a, "+" ,b ,"=" ,sum)
	fmt.Println(a, "-", b, "=" ,sub)

	f1 := addSub
	sum1, sub1 := f1(a, b)
	fmt.Println(a, "+" ,b ,"=" ,sum1)
	fmt.Println(a, "-", b, "=" ,sub1)
}