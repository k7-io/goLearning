package main

import (
	"fmt"
	"go_learning/cache"
)

var (
	db cache.DB
)

func init() {
	db = cache.DB{}
	db.Init()
}

func ErrHandle(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	size, err := db.Len("myLen")
	ErrHandle(err)
	fmt.Printf("size:%v\n", size)
}
