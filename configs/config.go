package configs

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var config Config

//配置
type Config struct {
	Env    string `yaml:"env"`
	JwtKey string `yaml:"jwt_key"`
	Db     Db     `yaml:"db"`
}

//数据配置
type Db struct {
	Host     string `yaml:"host"`
	Name     string `yaml:"name"`
	Password string `yaml:"password"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
}

func init() {
	//获取配置文件
	data, _ := ioutil.ReadFile("configs/config.yml")
	//把yaml形式的字符串解析成struct类型
	_ = yaml.Unmarshal(data, &config)
}

//获取配置文件
func GetConfig() *Config {
	return &config
}

//是否开发环境
func (c *Config) IsDev() bool {
	if c.Env == "dev" {
		return true
	}
	return false
}
