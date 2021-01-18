package main

import "fmt"

func main() {
	//logical operator
	a := true
	b := false
	fmt.Println("a&&b = ", a && b)
	fmt.Println("a||b = ", a || b)
	fmt.Println("!a = ", !a)
	//other operator
	c := 1
	var p *int
	p = &c
	fmt.Println("^c = ", ^c)
	fmt.Println("var a's address is ", p)
}