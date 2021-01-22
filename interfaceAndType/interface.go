package main

import (
	"errors"
	"fmt"
)

//interface define
type Animal interface {
	Name() string
	Speak() string
}

type Cat struct {

}

func (cat Cat) Name() string{
	return "Cat"
}

func (cat Cat) Speak() string {
	return "miao miao miao"
}

type IDatabaser interface {
	Connect() error
	Disconnect() error
}

type IRediser interface {
	Connect() error
}

//MySql database
type Mysql struct {
	DBName string
	isConnect bool
}

func (mysql *Mysql) Connect() error {
	fmt.Println("Mysql Connect DB => ", mysql.DBName)
	//do connect
	mysql.isConnect = true
	if mysql.isConnect {
		fmt.Println("Mysql connect success")
		return nil
	}else{
		return errors.New("Mysql connect failed")
	}
}

func (mysql *Mysql) Disconnect() error {
	//do disconnect
	fmt.Println("Mysql disconnect success")
	return nil
}

//Redis database
type Redis struct {
	DBName string
}

func (redis *Redis) Connect() error {
	fmt.Println("Redis Connect DB => ", redis.DBName)
	//do connect
	fmt.Println("Redis connect success")
	return nil
}

func (redis *Redis) Disconnect() error {
	//do disconnect
	fmt.Println("Redis disconnect success")
	return nil
}

//Interface oriented programming
func HandleDB(db IDatabaser)  {
	fmt.Println("start connect")
	db.Connect()
	fmt.Println("disconnect")
	db.Disconnect()
}

func main()  {
	var mysql = Mysql{"student", false}
	fmt.Println("start connect")
	mysql.Connect()
	//disconnect
	fmt.Println("disconnect")
	mysql.Disconnect()

	var redis = Redis{"teacher"}
	fmt.Println("start connect")
	redis.Connect()
	//disconnect
	fmt.Println("disconnect")
	redis.Disconnect()

	var mysql1 = Mysql{"student", false}
	HandleDB(&mysql1)
	var redis1 = Redis{"teacher"}
	HandleDB(&redis1)

	var redis2 = Redis{"teacher"}
	var idb IDatabaser = &redis2
	idb.Connect()
	idb.Disconnect()

	var idb1 IDatabaser = &Redis{"teacher"}
	var iredis IRediser
	iredis = idb1
	iredis.Connect()
}