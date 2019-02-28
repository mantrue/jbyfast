package setting

import (
	"log"
	"time"

	"gin-blog/pkg/logging"
	"github.com/go-ini/ini"
)

var (
	Cfg *ini.File

	RunMode string

	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	PageSize      int
	JwtSecret     string
	JWT_TIME_DIE  time.Duration
	MaxRequestNum float64 = 2
)

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

var RedisSetting = &Redis{}

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	LoadBase()
	LoadServer()
	LoadApp()
}

func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		logging.Info("Fail to get section 'server': %v" + err.Error())
	}

	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")

	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second

	redissec, err := Cfg.GetSection("redis")
	r := new(Redis)
	r.Host = redissec.Key("Host").MustString("")
	r.Password = redissec.Key("Password").MustString("")
	r.MaxIdle = redissec.Key("MaxIdle").MustInt(100)
	r.MaxActive = redissec.Key("MaxActive").MustInt(100)
	r.IdleTimeout = time.Duration(sec.Key("IdleTimeout").MustInt(60)) * time.Second
	RedisSetting = r
}

func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}

	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
	JWT_TIME_DIE = time.Duration(sec.Key("JWT_TIME_DIE").MustInt(60)) * time.Second
}
