package config

import (
	"github.com/jjggzz/kj/config"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var (
	Conf = &Config{}
)

type Config struct {
	Server    *config.Server    `yaml:"server"`
	Discovery *config.Discovery `yaml:"discovery"`
	DB        *config.DB        `yaml:"db"`
	Zipkin    struct {
		Address string `yaml:"address"`
	} `yaml:"zipkin"`
	GenerateKeyServerName string `yaml:"generateKeyServerName"`
}

func Init() {
	// 初始化配置
	FileContent, err := ioutil.ReadFile(config.ConfPath)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(FileContent, Conf)
	if err != nil {
		panic(err)
	}
}