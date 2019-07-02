package controllers

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"time"
)

type PublicController struct {
	Tstart    time.Time
	Tend      time.Time
	Tsub      time.Duration
	W         http.ResponseWriter
	R         *http.Request
	S         httprouter.Params
	TokenUser string
}

func (p *PublicController) ToJson(data interface{}) ([]byte, error) {
	bjson, err := json.Marshal(data)
	if err != nil {
		return []byte{}, err
	}
	return bjson, nil
}

func (p *PublicController) RpcSend() {

}
