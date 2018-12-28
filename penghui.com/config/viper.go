package config

import (
	"fmt"
	"github.com/spf13/viper"
)

const ConfigPath_URL = "E:/goroot/src/penghui.com/config"

var (
	Conf *Config
)

type Config struct {
	Mysql   MysqlConf
	Web     WebConf
	Runtime RuntimeConf
	Token   SecretConf
}

type MysqlConf struct {
	Debug    string `mapstructure:"debug"`
	Type     string `mapstructure:"type"`
	Hostname string `mapstructure:"hostname"`
	Database string `mapstructure:"database"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Hostport string `mapstructure:"hostport"`
	Dsn      string `mapstructure:"dsn"`
	Charset  string `mapstructure:"charset"`
	Prefix   string `mapstructure:"prefix"`
}

type WebConf struct {
	Debug string `mapstructure:"debug"`
}

type RuntimeConf struct {
	File string `mapstructure:"file"`
}

type SecretConf struct {
	Token string `mapstructure:"token"`
}

func init() {
	fmt.Println("config init...")
	BaseConfig()
}

func BaseConfig() error {
	Conf = NewConfigViper()
	viper.SetConfigName("base")         // 配置文件的名字
	viper.SetConfigType("toml")         // 配置文件的类型
	viper.AddConfigPath(ConfigPath_URL) // 配置文件的路径
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(&Conf); err != nil {
		panic(fmt.Errorf("unable to decode into struct：  %s \n", err))
	}

	return nil
}

func NewConfigViper() *Config {
	return &Config{}
}
