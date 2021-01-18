package main

import "fmt"

func pyramid(n int)  {
	//n is total number of layers
	for i := 1; i <= n; i++ {
		//print space first, the rule is the total number of layers minus the current number of layers
		for j := 1; j <= n - i; j++ {
			fmt.Print(" ")
		}
		//k is the number of * for per layer, the role is 2 * i - 1
		for k := 1; k <= 2 * i - 1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func main()  {
	n := 9
	pyramid(n)
}