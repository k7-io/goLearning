package main

import (
	"fmt"
	"os"
	"strconv"
)

func ReadFile(path string)  {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}else{
		buf := make([]byte, 1024)
		fmt.Println("The following is file content:")
		for {
			length, _ := file.Read(buf)
			if length == 0 {
				break
			}
			fmt.Println(string(buf))
		}
	}
	file.Close()
}

func ReadFileAt(path string)  {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}else {
		buf := make([]byte, 1024)
		fmt.Println("The following is file content:")
		_, _ = file.ReadAt(buf, 9)
		fmt.Println(string(buf))
		_ = file.Close()
	}
}

func main()  {
	file, err := os.OpenFile("/home/huoyinghui/test", os.O_RDWR|os.O_CREATE, 0766)
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println(file)
		file.Close()
	}

	//file read
	ReadFile("/home/huoyinghui/test")

	//file read at
	ReadFileAt("/home/huoyinghui/test")

	//file write in
	file1, err1 := os.Create("/home/huoyinghui/test")
	if err1 != nil {
		fmt.Println(err1)
	}else{
		data := "I am data\r\n"
		for i := 0; i < 3; i++{
			file1.Write([]byte(data))
		}
		file1.Close()
	}
	ReadFile("/home/huoyinghui/test")

	//file write in at
	file2, err2 := os.Create("/home/huoyinghui/test")
	if err2 != nil {
		fmt.Println(err2)
	}else{
		data := "I am data"
		for i := 0; i < 3; i++{
			ix := i * 64
			file2.WriteAt([]byte(data + strconv.Itoa(i) + "\r\n"), int64(ix))
		}
		file2.Close()
	}
	ReadFile("/home/huoyinghui/test")

	//delete file
	err3 := os.Remove("/home/huoyinghui/test1/test")
	if err3 != nil {
		fmt.Println(err3)
	}else{
		fmt.Println("Delete Success")
	}

	//delete all file
	err4 := os.RemoveAll("/home/huoyinghui/test1")
	if err4 != nil {
		fmt.Println(err4)
	}else{
		fmt.Println("Delete all success")
	}
}
