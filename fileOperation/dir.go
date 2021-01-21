package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func ListDir(dirPath string) error {
	dir, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return err
	}else{
		for _, fi := range dir {
			if fi.IsDir() == true {
				fmt.Println("dir:", fi.Name())
			}else{
				fmt.Println("file:", fi.Name())
			}
		}
	}
	return nil
}

func WalkDir(path string)  {
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error{
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		println(path)
		return nil
	})
	if err != nil {
		fmt.Printf("filePath.Walk() returned %v\n", err)
	}
}

func createDir(path, dirName string)  {
	dirPath := path + dirName
	fmt.Println("dirPath:", dirPath)
	err := os.Mkdir(dirPath, 0777)
	if err != nil {
		fmt.Println(err)
	}else{
		os.Chmod(dirPath, 0777)
		fmt.Println("Create dir => " + path + dirName)
	}
}

func createDirAll(path string, dirName string)  {
	dirPath := path + dirName
	fmt.Println("Create dir => " + dirPath)
	err := os.MkdirAll(dirPath, 0777)
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println("Create Success!")
		os.Chmod(dirPath, 0777)
	}
}

func deleteEmptyDir(dirPath string)  {
	fmt.Println("Delete dir => " + dirPath)
	err := os.Remove(dirPath)
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Println("Delete Success!")
	}
}

func deleteNotEmptyDir(dirPath string)  {
	fmt.Println("Delete dir => " + dirPath)
	err := os.RemoveAll(dirPath)
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println("Delete all Success!")
	}
}

func main()  {
	err := ListDir("/home/huoyinghui/go-dev/src")
	if err != nil {
		fmt.Println(err)
	}

	//Recursively traverse all the files in the folder
	//WalkDir("/home/huoyinghui/go-dev/src")

	//create dir:Mkdir
	createDir("/home/huoyinghui/", "test")

	//create dirs:MkdirAll
	createDirAll("/home/huoyinghui/", "dir1/dir2/dir3")

	//delete dir or file: delete empty dir
	deleteEmptyDir("/home/huoyinghui/test")
	deleteEmptyDir("/home/huoyinghui/test")

	//delete not empty dir
	deleteNotEmptyDir("/home/huoyinghui/dir1")
	deleteNotEmptyDir("/home/huoyinghui/dir1")
}
