package main

import "fmt"

type User struct{
	Name string
	Email string
}

//point tansfer
func (u *User) ChangeNamePoint()  {
	u.Name = "Tom"
}

//value transfer
func (u User) ChangeNameValue()  {
	u.Name = "Tom"
}

func main() {
	//point tansfer test
	u := User{
		"Peter",
		"go@go.com",
	}
	fmt.Println("Name:", u.Name, "Email:", u.Email)
	u.ChangeNamePoint()
	fmt.Println("Name:", u.Name, "Email:", u.Email)

	us := &User{
		"Peter",
		"go@go.com",
	}
	fmt.Println("Name:", us.Name, "Email:", us.Email)
	us.ChangeNamePoint()
	fmt.Println("Name:", us.Name, "Email:", us.Email)

	//value transfer test
	us1 := &User{
		"Peter",
		"go.go@com",
	}
	fmt.Println("Name:", us1.Name, "Email:", us1.Email)
	us1.ChangeNameValue()
	fmt.Println("Name:", us1.Name, "Email:", us1.Email)
}