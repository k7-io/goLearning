package main

import (
	"fmt"
	"reflect"
)

type Number int

type Person struct {

}

func checkType(t reflect.Type)  {
	if t.Kind() == reflect.Ptr {
		fmt.Printf("variable's type %v, point to variable",t.Kind())
		t = t.Elem()
	}
	fmt.Printf("variable's type name => %v, type's type => %v\n", t.Name(), t.Kind())
}

func checkValue(v reflect.Value)  {
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() == reflect.Int{
		var v1 int = int(v.Int())
		var v2 int = v.Interface().(int)
		fmt.Println(v1, v2)
	}
}

func Equal(a, b int) bool {
	if a == b {
		return true
	}
	return false
}

func main()  {
	var a interface{} = "I am string"
	typeOfa := reflect.TypeOf(a)
	fmt.Println("variable a's type is:", typeOfa.Name())
	valueOfa := reflect.ValueOf(a)
	if typeOfa.Kind() == reflect.String {
		fmt.Println("variable a's value is:", valueOfa.String())
	}

	var b int
	typeOfb := reflect.TypeOf(b)
	if typeOfb.Kind() == reflect.Int{
		fmt.Println("variable a's kind is int")
	}else{
		fmt.Println("variable a's kind is not int")
	}

	var num Number = 10
	typeOfNum := reflect.TypeOf(num)
	fmt.Println("typeOfNum")
	checkType(typeOfNum)
	var number int = 10
	fmt.Println("valueOfNum")
	valueOfNum := reflect.ValueOf(number)
	checkValue(valueOfNum)

	var person Person
	typeOfPerson := reflect.TypeOf(person)
	fmt.Println("typeOfPerson")
	checkType(typeOfPerson)
	typeOfPersonPtr := reflect.TypeOf(&person)
	fmt.Println("typeOfPersonPtr")
	checkType(typeOfPersonPtr)
	valueOfPersonPtr := reflect.ValueOf(&person)
	fmt.Println("valueOfPersonPtr")
	checkValue(valueOfPersonPtr)

	//Call
	valueOfFunc := reflect.ValueOf(Equal)
	args := []reflect.Value{reflect.ValueOf(1), reflect.ValueOf(2)}
	result := valueOfFunc.Call(args)
	fmt.Println("result:", result[0].Bool())
}