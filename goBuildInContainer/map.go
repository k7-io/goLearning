package main

import (
	"fmt"
	"sync"
)

var lock sync.RWMutex

func readMapLock(goMap map[int]int, key int) int {
	lock.Lock()
	m := goMap[key]
	lock.Unlock()
	return m
}

func writeMapLock(goMap map[int]int, key int, value int)  {
	lock.Lock()
	goMap[key] = value
	lock.Unlock()
}

func readMapSyncMap(goMap sync.Map, key int) int {
	res, ok := goMap.Load(key)
	if ok == true {
		return res.(int)
	}
	return 0
}

func writeMapSyncMap(goMap sync.Map, key int, value int)  {
	goMap.Store(key, value)
}

func main()  {
	//Initialize map
	//var studentScoreMap = map[string]int{
	//	"Tom":80,
	//	"Ben":90,
	//	"Peter":95,
	//}
	var studentScoreMap map[string]int
	studentScoreMap = make(map[string]int)
	studentScoreMap["Tom"] = 80
	studentScoreMap["Ben"] = 90
	studentScoreMap["Peter"] = 95
	fmt.Println("studentScoreMap's length:", len(studentScoreMap))
	fmt.Println(studentScoreMap)

	//Traverse map
	for k, v := range studentScoreMap{
		fmt.Println(k, v)
	}
	for k := range studentScoreMap{
		fmt.Println(k)
	}
	for _, v := range studentScoreMap{
		fmt.Println(v)
	}

	//delete map
	delete(studentScoreMap, "Tom")
	fmt.Println(studentScoreMap)

	//map concurrent operation
	//GoMap := make(map[int]int)
	var GoMap sync.Map
	for i := 0; i < 10000; i++ {
		//go readMapLock(GoMap, i)
		//go writeMapLock(GoMap, i, i)
		go readMapSyncMap(GoMap, i)
		go writeMapSyncMap(GoMap, i, i)
	}
	fmt.Println("Done")
}