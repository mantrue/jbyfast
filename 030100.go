package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	filepath.Walk("C:/Users/hack/Desktop/exerciseData0/030100", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && strings.Contains(path, ".txt") {
			_, fileurl := filepath.Split(path)
			res, _ := http.Get("http://www.phptest.me/appserver/get.php?code=030100&fileurl=" + fileurl)

			reson, _ := ioutil.ReadAll(res.Body)
			fmt.Println(string(reson))
		}
		return nil
	})
	//
}
