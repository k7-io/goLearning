package main

import "fmt"

func main()  {
	var a = 10
	fmt.Println("a=", a)
	a += 2
	fmt.Println("a += 2, a=", a)
	a -= 2
	fmt.Println("a -= 2, a=", a)
	a *= 2
	fmt.Println("a *= 2, a=", a)
	a /= 2
	fmt.Println("a /= 2, a=", a)
	a %= 2
	fmt.Println("a %= 2, a=", a)
}