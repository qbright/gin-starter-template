package config

import (
	"flag"
	"path/filepath"

	"github.com/spf13/viper"
)

var config *viper.Viper

func Init() {
	environment := flag.String("e", "development", "enviroment var ")
	flag.Parse()

	config = viper.New()
	config.Set("env", *environment)
	config.SetConfigName(*environment)
	config.AddConfigPath("../config/")
	config.AddConfigPath("config/")
	config.SetConfigType("yaml")
	config.AutomaticEnv()
	if err := config.ReadInConfig(); err != nil {
		panic(err)
	}
}

func relativePath(basedir string, path *string) {
	p := *path
	if len(p) > 0 && p[0] != '/' {
		*path = filepath.Join(basedir, p)
	}
}

func Env() string {
	return config.GetString("env")
}

func GetConfig() *viper.Viper {
	return config
}
