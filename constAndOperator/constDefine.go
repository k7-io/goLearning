package main

import (
	"fmt"
	"reflect"
)

func main() {
	const a float64= 3.1415
	const b = "hello world"
	fmt.Println(a)
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(b))
}