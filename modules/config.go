// Package modules
/*
   @author:xie
   @date:2022/5/24
   @note:
*/
package modules

import (
	"github.com/spf13/viper"
	"log"
	"time"
)

type Config struct {
	Mysql DataBase `yaml:"mysql"`
	Redis Cache    `yaml:"redis"`
	ES    Elastic  `yaml:"elastic"`
}
type DataBase struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
}
type Cache struct {
	Host        string        `yaml:"host"`
	Password    string        `yaml:"password"`
	Port        string        `yaml:"port"`
	DefaultDB   int           `yaml:"defaultDB"`
	DialTimeout time.Duration `yaml:"dialTimeout"`
}

type Elastic struct {
	Host  string `yaml:"host"`
	Port  string `yaml:"port"`
	Sniff bool   `yaml:"sniff"`
}

var Conf Config

func init() {
	cfg := viper.New()
	cfg.SetConfigName("config")
	cfg.SetConfigFile("./conf/app.yaml")
	if err := cfg.ReadInConfig(); err != nil { // 必须 先 读取 `ReadInConfig`
		log.Panicln(err)
	}
	err := cfg.Unmarshal(&Conf)
	log.Println(Conf)
	if err != nil {
		log.Panicln("参数配置失败")
	}
	log.Println("参数配置成功")
	cfg.WatchConfig()
}
