package controllers

import (
	"errors"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"penghui.com/ado"
	"penghui.com/lib/web"
	"runtime"
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
	_, err := ado.TechIndex()
	if err != nil {
		web.RespFail(w, r, 404, err, Filepath, Line)
	} else {
		web.RespJson(w, r, 200, map[string]string{"index": "index"}, Filepath, Line)
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
