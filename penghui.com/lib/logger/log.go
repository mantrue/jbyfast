package logger

import (
	"io"
	"log"
	"os"
)

var (
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

var ErrFileInfo *os.File
var err error

//var fileurl = config.Conf.Runtime.File

func init() {

	ErrFileInfo, err = os.OpenFile("E:/goroot/src/penghui.com/runtime/errors.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("打开日志文件失败：", err)
	}

	Info = log.New(io.MultiWriter(os.Stderr, ErrFileInfo), "Info:", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(io.MultiWriter(os.Stderr, ErrFileInfo), "Warning:", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(io.MultiWriter(os.Stderr, ErrFileInfo), "Error:", log.Ldate|log.Ltime|log.Lshortfile)

}

func TempText() string {
	return `[ {{ .SetTime }} ] 请求url: {{ .Request }}  请求方法:{{ .Method }} {{ .Url }}
[ info ] 当前请求:{{ .WebUrl }} [0s]
[ info ] [ 当前行号: {{ .Lline }} ]  当前文件: {{ .File }}
[ info ] [ 请求参数 ] [ {{ .Args }} ]`
}
