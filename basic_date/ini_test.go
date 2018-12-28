package basic_date

import (
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"testing"
)

func TestIni(t *testing.T) {
	viper.SetConfigName("conf")
	viper.SetConfigType("toml")
	viper.AddConfigPath("E:/goroot/src/basic_date/")
	if err := viper.ReadInConfig(); err != nil {
		t.Error(err.Error())
		t.Error("配置文件读取错误")
	}
	oks := viper.Get("base.pidfile")
	t.Logf("%s", oks)
}

func TestFile(t *testing.T) {
	file, err := os.Open("E:/goroot/src/basic_date/conf.toml")
	if err != nil {
		t.Error(err.Error())
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		t.Error(err.Error())
	}
	t.Logf("open file read")
	t.Logf("%s", data)
}
