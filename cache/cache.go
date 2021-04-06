package cache
/*
how to use:
	1.CMInit()
	2.c := GetCache
	3.c.Close()
*/

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

var (
	pool *redis.Pool
)

/*
server: 192.168.0.100:26379
password: ""
when len(password) more than 0, will do auth
*/
func newPool(server, password string) *redis.Pool {
	return &redis.Pool{
		MaxIdle: 3,
		IdleTimeout: 240 * time.Second,
		Dial: func () (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			if len(password) <= 0 {
				return c, err
			}
			if _, err := c.Do("AUTH", password); err != nil {
				c.Close()
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func CMInit(server, password string) {
	if pool == nil {
		pool = newPool(server, password)
	}
	if pool == nil {
		panic("cache pool init fail")
	}
}

func GetCache() redis.Conn{
	return pool.Get()
}
