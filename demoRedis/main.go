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

func (db *DB) InQueue(name string, values []interface{}) (err error) {
	size := len(values)
	for i := 0; i < size; i++ {
		_, err := db.client.Do("RPUSH", name, values[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func (db *DB) Len(name string) (int, error) {
	return redis.Int(db.client.Do("LLEN", name))
}

func (db *DB) OutQueueOne(name string) (value interface{}, err error) {
	reply, err := db.client.Do("LPOP", name)
	return reply, err
	//return db.client.Do("LPOP", name)
}

func (db *DB) OutQueue(name string) (values []interface{}, err error) {
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

func (db *DB) DelDBName(name string) (reply interface{}, err error) {
	return db.client.Do("del", name)
}

func StringSliceEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	if (a == nil) != (b == nil) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func main() {
	var db DB
	db.Init()
	db.Close()
	fmt.Println(db)
}
