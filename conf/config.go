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

type Mysql struct {
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

func InitConfig() {
	config := viper.New()
	config.SetConfigName("config")
	config.SetConfigFile(configPath)
	err := config.ReadInConfig()
	if err != nil {

	}
	config.WatchConfig()
}
