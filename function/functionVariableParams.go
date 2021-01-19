package main

import (
	"fmt"
)

func addSum(slice ... int) (sum int){
	sum = 0
	for _, v := range slice {
		sum += v
	}
	return sum
}

func addAll(slice []int)  {
	sum := 0
	for _, v := range slice {
		sum += v
	}
	fmt.Println(sum)
}

func add(num ... int)  {
	addAll(num)
}

func main()  {
	fmt.Println("1+2+3+4+5+6+7+8+9+10 =", addSum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}