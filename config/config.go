package config

import (
	"os"

	"github.com/spf13/viper"
)

var Config *viper.Viper

func Init() {
	path, err := os.Getwd()
	if err != nil {
		return
	}

	Config = viper.New()
	Config.AddConfigPath(path)
	Config.SetConfigName("config")
	Config.SetConfigType("yml")

	if err = Config.ReadInConfig(); err != nil {
		return
	}
}
