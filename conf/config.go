//Package conf
/*
   @author:xie
   @date:2022/5/24
   @note:
*/
package conf

import "github.com/spf13/viper"

type Mysql struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
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
