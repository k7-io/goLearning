package demoRedis

import (
	"testing"
)

var (
	db   DB
	name string
)

func init() {
	db = DB{}
	db.Init()
	name = "myTest"
}

func TestInQueue(t *testing.T) {
	_, err := db.InQueue(name, "1", "2", "3")
	if err != nil {
		t.Error(err)
	}
	db.Close()
	t.Log("TEST ", db)
}
