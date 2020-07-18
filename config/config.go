package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strconv"
)

// config.yaml配置项映射
type config struct {

	Server struct {
		ListenAddr string `yaml:"listen_addr"`
		Port int `yaml:"port"`
	} `yaml:"server"`

	DB struct {
		Driver  string `yaml:"driver"`
		User    string `yaml:"user"`
		Passwd  string `yaml:"passwd"`
		Host    string `yaml:"hosy"`
		Port    int	`yaml:"port"`
		DbName  string `yaml:"db_name"`
		Charset string `yaml:"charset"`
	} `yaml:"db"`
}

// 获取server监听地址
func (c config) GetServerListenAddr()string{
	return c.Server.ListenAddr + ":" + strconv.Itoa(c.Server.Port)
}

//
// 项目所需配置信息对象 单例
//
var c *config

//
// 获取config对象
//
func Config() *config  {
	if c == nil {
		configCtx,err := ioutil.ReadFile(`config/config.yaml`)
		if err != nil {
			panic("get config/config.yaml content error: " + err.Error())
		}

		err = yaml.Unmarshal(configCtx, &c)
		if err != nil {
			panic("yaml unmarshal failed error: " + err.Error())
		}
	}
	return c
}
