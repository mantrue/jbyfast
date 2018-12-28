package controllers

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"penghui.com/ado"
	"penghui.com/config"
	"penghui.com/lib/web"
	"runtime"
	"time"
)

var (
	Filepath string
	Line     int
)

func init() {
	_, Filepath, Line, _ = runtime.Caller(0)
}

type HomeController struct {
	PublicController
}

//首页
func (h *HomeController) MyLog(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	err := errors.New("")
	if err != nil {
		web.RespFail(w, r, 404, err, Filepath, Line)
	} else {
		web.RespJson(w, r, 200, map[string]string{"my": "data"}, Filepath, Line)
	}

}

//首页
func (h *HomeController) TechIndex(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	res, err := ado.UserAdo()
	if err != nil {
		web.RespFail(w, r, 404, err, Filepath, Line)
	} else {
		web.RespJson(w, r, 200, res, Filepath, Line)
	}

}

//新闻首页
func (h *HomeController) NewsIndex(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	strByte, err := h.ToJson(map[string]string{"hello": "1"})
	if len(strByte) > 0 {
		err = errors.New("not find")
		web.RespFail(w, r, 404, err, Filepath, Line)
	} else {
		web.RespJson(w, r, 200, map[string]string{"news": "news"}, Filepath, Line)
	}

}

//登录
func (h *HomeController) Login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = "Jon Snow"
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	t, err := token.SignedString([]byte(config.Conf.Token.Token))

	if err != nil {
		web.RespFail(w, r, 500, err, Filepath, Line)
	}

	web.RespJson(w, r, 200, map[string]string{"token": t}, Filepath, Line)
}
