package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var conf = viper.New()

func initConfig() {

	configPath := pflag.String("config", "", "config file path")
	pflag.Parse()
	conf.SetConfigType("toml")
	conf.SetDefault("version", "v1.0.0")
	if *configPath == "" {
		conf.SetConfigName("config")
		conf.AddConfigPath("/etc/go-realtimechat/")
		conf.AddConfigPath("$HOME/.go-realtimechat")
		conf.AddConfigPath(".")
	} else {
		conf.SetConfigFile(*configPath)
	}
	err := conf.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	file := conf.ConfigFileUsed()
	if file != "" {
		fmt.Println("Use config file: " + file)
	}
}
