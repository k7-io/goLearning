package main

import "fmt"

func Log(name string, i interface{})  {
	fmt.Printf("Name = %s, Type = %T, value = %v\n", name, i, i)
}

func Log1(args ... interface{})  {
	for num, arg := range args{
		fmt.Printf("Index => %d, value => %v\n", num, arg)
	}
}

func main()  {
	var v1 interface{} = 1
	var v2 interface{} = "abc"
	var v3 interface{} = true
	var v4 interface{} = &v1
	var v5 interface{} = struct {
		Name string
	}{"Xiao Ming"}
	var v6 interface{} = &v5
	Log("v1", v1)
	Log("v2", v2)
	Log("v3", v3)
	Log("v4", v4)
	Log("v5", v5)
	Log("v6", v6)

	//take value from empty intreface
	var a string = "abc"
	var i interface{} = a
	var b string = i.(string)
	fmt.Println(b)

	s := make([]interface{}, 3)
	s[0] = 1
	s[1] = "abc"
	s[2] = struct {
		Num int
	}{1}
	fmt.Println("variable length slice:")
	Log1(s...)
	fmt.Println("directly use slice:")
	Log1(s)
}