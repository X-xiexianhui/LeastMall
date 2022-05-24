//Package conf
/*
   @author:xie
   @date:2022/5/24
   @note:
*/
package conf

import (
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

var configPath = "./app.yaml"
var Conf *Config

func init() {
	viper.SetConfigName("config")
	viper.SetConfigFile(configPath)
	err := viper.Unmarshal(Conf)
	if err != nil {

	}
	viper.WatchConfig()
}
