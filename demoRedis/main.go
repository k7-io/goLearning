package demoRedis

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

type DB struct {
	client redis.Conn
}

func (db *DB) Init() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		// handle error
		panic("redis conn err:" + err.Error())
	}
	db.client = c
	//defer c.Close()
}

func (db *DB) Close() {
	if db.client != nil {
		db.client.Close()
	}
}

func (db *DB) InQueue(name string, values ...string) (reply interface{}, err error) {
	return db.client.Do("RPUSH", name, values)
}

func (db *DB) Len(name string) (int, error) {
	return redis.Int(db.client.Do("LLEN", name))
}

func (db *DB) OutQueueOne(name string) (value string, err error) {
	return redis.String(db.client.Do("LPOP", name))
}

func (db *DB) OutQueue(name string) (values []string, err error) {
	size, err := db.Len(name)
	if err != nil {
		return nil, err
	}
	for i := 0; i < size; i++ {
		ele, err := db.OutQueueOne(name)
		if err != nil {
			return nil, err
		}
		values = append(values, ele)
	}
	return values, nil
}

func main() {
	var db DB
	db.Init()
	db.Close()
	fmt.Println(db)
}
