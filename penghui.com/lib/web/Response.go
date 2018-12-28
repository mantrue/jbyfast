package web

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog"
	"io/ioutil"
	"net/http"
	"os"
	"penghui.com/config"
	"penghui.com/lib/logger"
	"strconv"
	"strings"
	"text/template"
	"time"
)

var (
	HTTP_RIGHT_REQUEST_NUM int64 = 0
	HTTP_FAIL_REQUEST_NUM  int64 = 0
)

//返回错误的res
func RespFail(w http.ResponseWriter, r *http.Request, status int, err error, file string, line int) {
	w.Header().Set("Connection", "close")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err != nil {
		res := map[string]string{"error": err.Error()}
		json.NewEncoder(w).Encode(res)
	}
	r.ParseForm()

	//记录log日志
	HTTP_FAIL_REQUEST_NUM++

	log := zerolog.New(os.Stderr).With().Timestamp().Logger()
	zerolog.TimeFieldFormat = "2006-01-02 15:04:05"

	if config.Conf.Web.Debug == "true" {
		fmt.Println("fail request", HTTP_FAIL_REQUEST_NUM)
		//记录请求日志
		log.Info().Msg(err.Error())

		var args string
		if r.Method == http.MethodGet {
			args = r.URL.RawQuery
		} else {
			body_byte, _ := ioutil.ReadAll(r.Body)
			body := string(body_byte)
			body = strings.Replace(body, "\r\n\r\n", "", -1)
			body = strings.Replace(body, "\r\n", "", -1)
			args = body
		}

		errdata := map[string]interface{}{"Args": args, "Lline": line, "File": file, "SetTime": time.Now(), "error": err.Error(), "Request": r.URL, "Ip": r.Host, "Method": r.Method, "Url": r.URL.Path, "WebUrl": r.Host + r.URL.Path, "Line": "a.php"}
		text := logger.TempText()
		tmpl, err := template.New("test").Parse(text + "\n\r")
		if err != nil {
			panic(err)
		}

		tmpl.Execute(logger.ErrFileInfo, errdata)
	} else {
		//只记录日志
		var args string
		if r.Method == http.MethodGet {
			args = r.URL.RawQuery
		} else {
			body_byte, _ := ioutil.ReadAll(r.Body)
			body := string(body_byte)
			body = strings.Replace(body, "\r\n\r\n", "", -1)
			body = strings.Replace(body, "\r\n", "", -1)
			args = body
		}

		errdata := map[string]interface{}{"Args": args, "Lline": line, "File": file, "SetTime": time.Now(), "error": err.Error(), "Request": r.URL, "Ip": r.Host, "Method": r.Method, "Url": r.URL.Path, "WebUrl": r.Host + r.URL.Path, "Line": "a.php"}
		text := logger.TempText()
		tmpl, err := template.New("test").Parse(text + "\n\r")
		if err != nil {
			panic(err)
		}

		tmpl.Execute(logger.ErrFileInfo, errdata)
	}
}

//返回正确的res
func RespJson(w http.ResponseWriter, r *http.Request, status int, value interface{}, file string, line int) (int, error) {
	data, err := json.Marshal(value)
	if err != nil {
		return 0, err
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(status)
	w.Write(data)

	r.ParseForm()

	//记录log日志
	HTTP_FAIL_REQUEST_NUM++

	log := zerolog.New(os.Stderr).With().Timestamp().Logger()
	zerolog.TimeFieldFormat = "2006-01-02 15:04:05"

	if config.Conf.Web.Debug == "true" {
		fmt.Println("fail request", HTTP_FAIL_REQUEST_NUM)

		var args string
		if r.Method == http.MethodGet {
			args = r.URL.RawQuery
		} else {
			body_byte, _ := ioutil.ReadAll(r.Body)
			body := string(body_byte)
			body = strings.Replace(body, "\r\n\r\n", "", -1)
			body = strings.Replace(body, "\r\n", "", -1)
			args = body
		}
		//记录请求日志
		log.Info().Msg("weburl:" + r.Host + r.URL.Path + "=====args:" + args)

		errdata := map[string]interface{}{"Args": args, "Lline": line, "File": file, "SetTime": time.Now(), "error": "", "Request": r.URL, "Ip": r.Host, "Method": r.Method, "Url": r.URL.Path, "WebUrl": r.Host + r.URL.Path, "Line": "a.php"}
		text := logger.TempText()
		tmpl, err := template.New("test").Parse(text + "\n\r")
		if err != nil {
			panic(err)
		}

		tmpl.Execute(logger.ErrFileInfo, errdata)
	} else {
		//只记录日志
		var args string
		if r.Method == http.MethodGet {
			args = r.URL.RawQuery
		} else {
			body_byte, _ := ioutil.ReadAll(r.Body)
			body := string(body_byte)
			body = strings.Replace(body, "\r\n\r\n", "", -1)
			body = strings.Replace(body, "\r\n", "", -1)
			args = body
		}

		errdata := map[string]interface{}{"Args": args, "Lline": line, "File": file, "SetTime": time.Now(), "error": "", "Request": r.URL, "Ip": r.Host, "Method": r.Method, "Url": r.URL.Path, "WebUrl": r.Host + r.URL.Path, "Line": "a.php"}
		text := logger.TempText()
		tmpl, err := template.New("test").Parse(text + "\n\r")
		if err != nil {
			panic(err)
		}

		tmpl.Execute(logger.ErrFileInfo, errdata)
	}

	return len(data), err
}
