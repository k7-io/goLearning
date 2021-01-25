package main

import (
	"fmt"
	"reflect"
)

type Person2 struct {
	Name string
	age int `json:"age"`
	string
}

func main()  {
	var a int = 5
	//interface to reflect.Type and reflect.Value
	fmt.Println(reflect.TypeOf(a))
	fmt.Printf("type:%T\n", reflect.TypeOf(a))
	fmt.Println(reflect.ValueOf(a))
	fmt.Printf("value:%T\n", reflect.ValueOf(a))

	//reflect.Value to interface
	valueOfa := reflect.ValueOf(a)
	fmt.Println(valueOfa.Interface())

	//CanSet
	person := Person2{"xiao zhang", 10, "notes"}
	valueOfPerson := reflect.ValueOf(&person)
	typeOfPerson := reflect.TypeOf(&person)
	for i := 0; i < valueOfPerson.Elem().NumField(); i++ {
		fieldValue := valueOfPerson.Elem().Field(i)
		fieldType := typeOfPerson.Elem().Field(i)
		fmt.Printf("name:%v, canAddress:%v, canSet:%v \n", fieldType.Name, fieldValue.CanAddr(), fieldValue.CanSet())
	}
	fmt.Println("before change:", person)
	valueOfPerson.Elem().Field(0).SetString("XIAO MING")
	fmt.Println("After change:", person)
}
