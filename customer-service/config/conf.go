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
	Server             *config.Server    `yaml:"server"`
	Discovery          *config.Discovery `yaml:"discovery"`
	DB                 *config.DB        `yaml:"db"`
	Zipkin             *config.Zipkin    `yaml:"zipkin"`
	Redis              *config.Redis     `yaml:"redis"`
	SecretCode         string            `yaml:"secretCode"`
	GenerateServerName string            `yaml:"generateServerName"`
	VerifyServerName   string            `yaml:"verifyServerName"`
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
