package main

//func add(x int,y int) (sum int) {
//	sum = x + y
//	return sum
//}

//func add(x, y int) (sum int)  {
//	sum = x + y
//	return sum
//}

func returnValue() (int, int)  {//not recommended
	return 0, 1
}

//func defaultValue() (a int, b string, c bool)  {//return 0, "", false
//	return
//}

func defaultValue() (a int, b string, c bool)  {
	return 1, "hello", true
}