package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"os"
	"penghui.com/config"
	"penghui.com/lib/logger"
)

var (
	engine  *xorm.Engine
	err     error
	fileurl = config.Conf.Runtime.File
)

func init() {
	var err error
	engine, err = xorm.NewEngine("mysql", "root:Jby&*2016@tcp(118.190.65.33:3306)/jingtongcloud?charset=utf8")
	if err != nil {
		logger.Error.Println(err.Error())
	}
	engine.ShowSQL(true)
	engine.Logger().SetLevel(core.LOG_DEBUG)
	f, err := os.OpenFile(fileurl, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		println(err.Error())
		return
	}
	engine.SetLogger(xorm.NewSimpleLogger(f))
}

type User struct {
	Id        int64  `json:"id"`
	Auth_name string `json:"auth_name"`
}

func UserModel() ([]User, error) {

	var results []User
	err := engine.Table("account_auth").Find(&results)
	if err != nil {
		panic(err)
	}
	return results, nil
}
