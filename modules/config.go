//Package conf
/*
   @author:xie
   @date:2022/5/24
   @note:
*/
package modules

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	Mysql DataBase `yaml:"mysql"`
	Redis Cache    `yaml:"redis"`
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
	DefaultDB   int           `yaml:"defaultDB"`
	DialTimeout time.Duration `yaml:"dialTimeout"`
}

var configPath = "../conf/app.yaml"
var Conf *Config

func init() {
	viper.SetConfigName("config")
	viper.SetConfigFile(configPath)
	err := viper.Unmarshal(Conf)
	if err != nil {
		logrus.Errorln("参数配置失败")
	}
	viper.WatchConfig()
	logrus.Println("参数配置成功")
}
