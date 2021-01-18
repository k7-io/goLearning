package main

import "fmt"

//1 1 2 3 5 8 13 21 34 ...
//func fibonacci(n int) (res int)  {
//	a := 1
//	b := 1
//	for i := 3; i <= n; i++ {
//		c := b
//		b = a + b
//		a = c
//	}
//	return b
//}

//func fibonacci(n int) (res int)  {
//	if n == 1 || n == 2 {
//		return 1
//	}
//	return fibonacci(n - 1) + fibonacci(n - 2)
//}

func fibonacci(n int) (res int)  {
	switch n{
	case 1:
		return 1
	case 2:
		return 1
	default:
		return fibonacci(n - 1) + fibonacci(n - 2)
	}
}
func main()  {
	n := 3
	fmt.Printf("The %d's value is %d", n, fibonacci(n))
}