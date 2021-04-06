package cache

import (
	"fmt"
	"goLearning/config"
	"testing"
)

func init() {
	appConf, err = config.LoadConf("../conf/app.yml")
	if err != nil {
		panic(err)
	}
	fmt.Printf("RedisConf:%v\n", RedisConf)
	CMInit(appConf.RedisConf, appConf.RedisConfPasswd)
}


func TestCMClient(t *testing.T) {
	CMInit(appConf.RedisConf, appConf.RedisConfPasswd)
	c := GetCache()
	if c == nil {
		t.Fail()
	}
	defer c.Close()
	t.Logf("CM pool:%v\n", c)
}