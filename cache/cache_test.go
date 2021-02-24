package cache

import (
	"testing"
)

var (
	db DB
)

func init() {
	db = DB{}
	db.Init()
}

func TestInQueue(t *testing.T) {
	var dataTests = []struct {
		inData   []string
		expected error
	}{
		{
			inData:   []string{"1", "2", "3"},
			expected: nil,
		},
	}
	for _, dt := range dataTests {
		name := "myTestInQueue"
		db.client.Do("del", name)
		inData := make([]interface{}, len(dt.inData))
		for i := 0; i < len(inData); i++ {
			inData[i] = dt.inData[i]
		}
		err := db.InQueue(name, inData)
		if dt.expected != err {
			t.Errorf("TestInQueue inData:%v err:%v", dt.inData, err)
		}
	}

	db.Close()
	t.Log("TEST ", db)
}

func TestLen(t *testing.T) {
	var dataTests = []struct {
		inData   []int
		expected int
	}{
		{
			inData:   []int{1, 2, 3},
			expected: 3,
		},
	}
	for _, dt := range dataTests {
		name := "myLen"
		db.client.Do("del", name)
		inData := make([]interface{}, len(dt.inData))
		for i := 0; i < len(inData); i++ {
			inData[i] = dt.inData[i]
		}
		err := db.InQueue(name, inData)
		if err != nil {
			t.Errorf("TestLen inQueue inData:%v err:%v", dt.inData, err)
		}
		size, err := db.Len(name)
		if err != nil {
			t.Errorf(" TestLen len inData:%v err:%v", dt.inData, err)
		}
		if size != dt.expected {
			t.Errorf("TestLen size:%v dt.expected:%v inData:%v err:%v",
				size, dt.expected, dt.inData, err)
		}
	}

	db.Close()
	t.Log("TEST ", db)
}

func TestOutQueue(t *testing.T) {
	var dataTests = []struct {
		inData   []string
		expected []string
	}{
		{
			inData:   []string{"1", "2", "3"},
			expected: []string{"1", "2", "3"},
		},
	}
	for _, dt := range dataTests {
		name := "myTestOutQueue"
		db.client.Do("del", name)
		inData := make([]interface{}, len(dt.inData))
		for i := 0; i < len(inData); i++ {
			inData[i] = dt.inData[i]
		}
		err := db.InQueue(name, inData)
		if err != nil {
			t.Errorf("TestOutQueue inQueue inData:%v err:%v", dt.inData, err)
		}
		res, err := db.OutQueue(name)
		if err != nil {
			t.Errorf("TestOutQueue outQueue err:%v", err)
		}
		var resStr []string
		for i := 0; i < len(res); i++ {
			resStr = append(resStr, string(res[i].([]uint8)))
		}
		if StringSliceEqual(resStr, dt.expected) != true {
			t.Errorf("outQueue res:%v dt.expected:%v",
				res, dt.expected)
		}
	}

	db.Close()
	t.Log("TEST ", db)
}
