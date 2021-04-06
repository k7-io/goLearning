package cache

import (
	"github.com/gomodule/redigo/redis"
)

type Queue interface {
	Len(name string) (int, error)
	InQueue(name string, values []interface{}) (err error)
	OutQueueOne(name string) (value interface{}, err error)
	DelKey(name string) (reply interface{}, err error)
}


type DBQueue struct {
	Client redis.Conn
}


func (db *DBQueue) Init(conn redis.Conn) {
	if conn == nil {
		conn = GetCache()
	}
	db.Client = conn
}

func (db *DBQueue) Close() {
	if db.Client != nil {
		db.Client.Close()
	}
}
func (db *DBQueue) Check() {
	if db.Client == nil {
		panic("db client must be init")
	}
}

func (db *DBQueue) InQueue(name string, values []interface{}) (err error) {
	db.Check()
	size := len(values)
	for i := 0; i < size; i++ {
		_, err := db.Client.Do("RPUSH", name, values[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func (db *DBQueue) Len(name string) (int, error) {
	db.Check()
	return redis.Int(db.Client.Do("LLEN", name))
}

func (db *DBQueue) OutQueueOne(name string) (value interface{}, err error) {
	db.Check()
	reply, err := db.Client.Do("LPOP", name)
	return reply, err
}

func (db *DBQueue) OutQueue(name string) (values []interface{}, err error) {
	db.Check()
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

func (db *DBQueue) DelKey(name string) (reply interface{}, err error) {
	db.Check()
	return db.Client.Do("del", name)
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
