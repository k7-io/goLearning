package main

import (
	"fmt"
	"reflect"
)

type Person1 struct {
	Name string
	Age int `json:"age"`
	string
}

func (p Person1) GetName()  {
	fmt.Println(p.Name)
}

func main()  {
	person := Person1{"xiao zhang", 10, "notes"}
	//reflect.TypeOf
	typeOfPerson := reflect.TypeOf(person)
	fmt.Println("traverse struct")
	for i := 0; i < typeOfPerson.NumField(); i++ {
		field := typeOfPerson.Field(i)
		fmt.Printf("name:%v, tag:%v, anonymous:%v \n", field.Name, field.Tag, field.Anonymous)
	}
	if field, ok := typeOfPerson.FieldByName("Age"); ok {
		fmt.Println("according to name")
		fmt.Printf("name:%v, tag json:%v\n", field.Name, field.Tag.Get("json"))
	}
	field := typeOfPerson.FieldByIndex([]int{1})
	fmt.Println("according to index")
	fmt.Printf("name:%v, tag:%v\n", field.Name, field.Tag)

	//reflect.ValueOf
	valueOfPerson := reflect.ValueOf(person)
	fmt.Printf("person field nums: %d\n",valueOfPerson.NumField())
	fmt.Println("Field")
	field1 := valueOfPerson.Field(1)
	fmt.Printf("field value:%v\n", field1.Int())
	field2 := valueOfPerson.FieldByName("Age")
	fmt.Println("FieldByName")
	fmt.Printf("field value:%v\n", field2.Interface())
	field3 := valueOfPerson.FieldByIndex([]int{0})
	fmt.Println("FieldByIndex")
	fmt.Printf("field value:%v\n", field3.String())

	f := valueOfPerson.MethodByName("GetName")
	fmt.Println(f)
	f.Call([]reflect.Value{})
}
