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
	Zipkin    *config.Zipkin    `yaml:"zipkin"`
	Redis     struct {
		Address     string `yaml:"address"`
		MaxIdle     int    `yaml:"maxIdle"`
		MaxActive   int    `yaml:"maxActive"`
		IdleTimeout int64  `yaml:"idleTimeout"`
	} `yaml:"redis"`
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
