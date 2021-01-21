package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"reflect"
	"runtime"
)

func protect()  {
	defer func(){
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	panic("Serious bug")
}

func protectFunc(f func())  {
	defer func(){
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	f()
}

func div(dividend int, divisor int) (int, error)  {
	if divisor == 0 {
		return 0, errors.New("divisor is zero")
	}
	return dividend/divisor, nil
}

type ErrNegSqrt float64

func (e ErrNegSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func sqrt(x float64) (float64, error)  {
	if x < 0.0 {
		fmt.Println("ErrNegSqrt(x) type:::",reflect.TypeOf(ErrNegSqrt(x)))
		return 0, ErrNegSqrt(x)
	}else{
		return math.Sqrt(x), nil
	}
}

func main()  {
	//error handling
	f, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println(f)
	}

	//custom error
	err1 := errors.New("this is an error")
	var err2 error
	fmt.Println(err1.Error())
	fmt.Println(err2)

	//custom error format
	if _, _, line, ok := runtime.Caller(0);ok == true{
		err := fmt.Errorf("***line %d error***", line)
		fmt.Println(err.Error())
	}

	//go language down and recover
	//panic("Serious bug.")
	//fmt.Println("Invalid code")
	defer func(){
		fmt.Println("func main exit")
	}()
	protect()
	fmt.Println("valid code")

	//recover practical application
	protectFunc(func(){
		fmt.Println("This is function 1")
		panic("Serios bug from function 1")
	})
	protectFunc(func(){
		fmt.Println("This is function 2")
		panic("Serios bug from function 2")
	})
	fmt.Println("Valid code")

	//go error application,panic recover not recommended,return value recommended
	res1, err3 := div(4, 2)
	if err3 != nil {
		fmt.Println(err3.Error())
	}else{
		fmt.Println("4/2 = ", res1)
	}
	res2, err4 := div(4, 0)
	if err4 != nil {
		fmt.Println(err4.Error())
	}else{
		fmt.Println("4/0 = ", res2)
	}

	//error interface application
	res, err5 := sqrt(-2)
	if err5 != nil {
		fmt.Println(err5.Error())
	}else{
		fmt.Println(res)
	}

	type alis int
	n := 2
	fmt.Println(reflect.TypeOf(alis(n)))
}