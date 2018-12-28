package controllers

import (
	"encoding/json"
	"time"
)

type PublicController struct {
	Tstart time.Time
	Tend   time.Time
	Tsub   time.Duration
}

func (p *PublicController) ToJson(data interface{}) ([]byte, error) {
	bjson, err := json.Marshal(data)
	if err != nil {
		return []byte{}, err
	}
	return bjson, nil
}
