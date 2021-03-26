
package config

import (
	"encoding/json"
	"errors"
	"gopkg.in/yaml.v2"
	iot "io/ioutil"
	"strings"
)

type AppConfInfo struct {
	AppName         string `yaml:"app_name" json:"app_name"`
	Host        	int    `yaml:"host" json:"host"`
	HttpPort        int    `yaml:"http_port" json:"http_port"`
	RunMode         string `yaml:"run_mode" json:"run_mode"`
	PgDataSource    string `yaml:"pg_data_source" json:"pg_data_source"`
	OrmDebug        bool   `yaml:"orm_debug" json:"orm_debug"`
	EnableDocs      bool   `yaml:"enable_docs" json:"enable_docs"`
	LogLevel        int    `yaml:"log_level" json:"log_level"`
	JwtSalt         string `yaml:"jwt_salt" json:"jwt_salt"`
	RedisConf       string `yaml:"redis_conf" json:"redis_conf"`
	RedisConfPasswd string `yaml:"redis_conf_passwd" json:"redis_conf_passwd"`
	OSSConf 		OSSConfInfo `yaml:"oss_conf" json:"oss_conf"`
	DocConf			DocConfInfo `yaml:"doc_conf" json:"doc_conf"`
}
type DocConfInfo struct {
	ApiPre       	string    `yaml:"api_pre" json:"api_pre"`
	StaticFilePath  string    `yaml:"static_file_path" json:"static_file_path"`
}

type OSSConfInfo struct {
	EndPoint         string `yaml:"end_point" json:"end_point"`
	Bucket    		 string `yaml:"bucket" json:"bucket"`
	AccessKeyID      string `yaml:"access_key_id" json:"access_key_id"`
	AccessKeySecret  string `yaml:"access_key_secret" json:"access_key_secret"`
}

func LoadConf(filepath string) (item *AppConfInfo, err error) {
	if filepath == "" {
		return nil, errors.New("filepath is empty, must use --config xxx.yml/json")
	}

	data, err := iot.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	item = &AppConfInfo{}
	if strings.HasSuffix(filepath, ".json") {
		err = json.Unmarshal(data, item)
	} else if strings.HasSuffix(filepath, ".yml") || strings.HasSuffix(filepath, ".yaml") {
		err = yaml.Unmarshal(data, item)
	} else {
		return nil, errors.New("you config file must be json/yml")
	}

	if err != nil {
		return nil, err
	}

	return item, nil
}

func LoadConfFromData(data []byte, t string) (item *AppConfInfo, err error) {
	// data 可以是json/yml格式的数据. 调用方需要指定t.
	item = &AppConfInfo{}
	if t == "json" {
		err = json.Unmarshal(data, item)
	} else if t == "yml" {
		err = yaml.Unmarshal(data, item)
	}
	if err != nil {
		return nil, err
	}

	return item, nil
}