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
	Email     struct {
		QQ struct {
			Sender   string `yaml:"sender"`
			Password string `yaml:"password"`
			Addr     string `yaml:"addr"`
			Port     int    `yaml:"port"`
		} `yaml:"qq"`
	} `yaml:"email"`
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
