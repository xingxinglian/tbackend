package config

import (
	"flag"
	"fmt"
	"os"

	"testing"

	"gopkg.in/yaml.v2"
)

var Config Conf

type Conf struct {
	Env     string `yaml:"env"`
	AppName string `yaml:"app_name"`

	Aws struct {
		S3Region    string `yaml:"s3_region"`
		S3Bucket    string `yaml:"s3_bucket"`
		S3KeyId     string `yaml:"s3_key_id"`
		S3AccessKey string `yaml:"s3_access_key"`
	} `yaml:"aws"`
}

var gConfigName string

func init() {
	testing.Init()
	fmt.Println("config init")
	flag.StringVar(&gConfigName, "conf", "./conf.yaml", "config name")
	flag.Parse()

	fmt.Println("config name ", gConfigName)
	ParseYaml(gConfigName, &Config)
}

func ParseYaml(file string, configRaw interface{}) {
	content, err := os.ReadFile(file)
	if err != nil {
		panic("加载配置文件错误" + file + "错误原因" + err.Error())
	}

	err = yaml.Unmarshal(content, configRaw)
	if err != nil {
		panic("解析配置文件错误" + file + "错误原因" + err.Error())
	}
}
