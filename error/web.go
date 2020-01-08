package main

import (
	"error/yewu"
	"net/http"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func errWrapper(handeler appHandler) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		err := handeler(writer, request)
		if err != nil {

		}
	}
}

func main() {
	http.HandleFunc("/list", yewu.List)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
