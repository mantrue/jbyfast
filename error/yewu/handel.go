package yewu

import (
	"io/ioutil"
	"net/http"
	"os"
)

func List(writer http.ResponseWriter, request *http.Request) {
	path := request.URL.Path
	file, err := os.Open(path)
	if err != nil {
		http.Error(writer, err.Error(), 500)
	}
	defer file.Close()
	all, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	writer.Write(all)
}
