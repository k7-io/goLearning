package test

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	 "goLearning/config"
	"testing"
)
var (
	appConf *config.AppConfInfo
	ossConf config.OSSConfInfo
	bucket *oss.Bucket
	err error
)

func init() {
	appConf, err = config.LoadConf("../conf/app.yml")
	if err != nil {
		panic(err)
	}
	ossConf = appConf.OSSConf
	fmt.Printf("ossConf:%v\n", ossConf)
	client, err := oss.New(ossConf.EndPoint, ossConf.AccessKeyID, ossConf.AccessKeySecret)
	if err != nil {
		panic(err)
	}
	// 获取存储空间。
	bucket, err = client.Bucket(ossConf.Bucket)
	if err != nil {
		panic(err)
	}
	fmt.Printf("bucket:%v\n", bucket)
}

func handleError(err error, t *testing.T) {
	t.Error(err)
	t.Fail()
}

func TestPutObjectFromFile(t *testing.T) {
	err := bucket.PutObjectFromFile("test01.txt","./a.txt")
	if err != nil {
		handleError(err, t)
	}
}

func TestGetObjectToFile(t *testing.T) {
	err := bucket.GetObjectToFile("test.txt","./da_tmp.txt")
	if err != nil {
		handleError(err, t)
	}
}
