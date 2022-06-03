// Package conn
/*
   @author:xie
   @date:2022/5/24
   @note:
*/
package conn

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"time"
)

type Config struct {
	Mysql   DataBase `yaml:"mysql"`
	Redis   Cache    `yaml:"redis"`
	Elastic es       `yaml:"elastic"`
	Domain  string   `yaml:"domain"`
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

type es struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
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
	if err != nil {
		log.Panicln("参数配置失败")
	}
	log.Println("参数配置成功")
	cfg.WatchConfig()
	fmt.Println(Conf)
}
