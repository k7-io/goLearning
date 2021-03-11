package config

import (
	"gopkg.in/yaml.v2"
	iot "io/ioutil"
	"testing"
)


var TestData = []byte(`
appname: bdev
httpport: 5008
runmode: prod
pgdatasource: user=postgres password=postgres dbname=test host=127.0.0.1 port=5432
  sslmode=disable
ormdebug: true
enabledocs: true
loglevel: 7
jwt_salt: testsalt
oss_conf:
  end_point: http://oss-cn-beijing.aliyuncs.com
  access_key_id: xxxxx
  access_key_secret: xxxxx

`)


func TestNewConf(t *testing.T) {
	conf := AppConfInfo{
		AppName:         "bdev",
		HttpPort:        5008,
		RunMode:         "prod",
		EnableDocs:      true,
		OrmDebug:        true,
		PgDataSource:    "user=postgres password=postgres dbname=test host=127.0.0.1 port=5432 sslmode=disable",
		JwtSalt:         "testsalt",
		RedisConf: 		 "127.0.0.1:6379",
		OSSConf: OSSConfInfo{
			EndPoint: "http://oss-cn-beijing.aliyuncs.com",
			AccessKeyID: "xxxxx",
			AccessKeySecret: "xxxxx",
			Bucket: "test",
		},
		DocConf: DocConfInfo{
			ApiPre: "http://www.hyh.com:8000",
			StaticFilePath: "./static",
		},
	}

	data, err := yaml.Marshal(conf)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("data:%v", string(data))
	if err := iot.WriteFile("../conf/app_template.yml", data, 0644); err != nil {
		t.Fatal(err)
	}

}

func TestLoadConf(t *testing.T) {
	data, err := LoadConf("../conf/app.yml")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("data:%v", data)
}

func TestLoadConfFromData(t *testing.T) {
	data, err := LoadConfFromData(TestData, "yml")
	if err != nil {
		t.Logf("data:%v", data)
		t.Fatal(err)
	}
	t.Logf("data:%v", data)
}