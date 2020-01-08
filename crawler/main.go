package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error:status code ", resp.StatusCode)
		return
	} else {
		buf := bufio.NewReader(resp.Body)
		e := determineEncoding(buf)
		utf8Reader := transform.NewReader(buf, e.NewDecoder())
		all, err := ioutil.ReadAll(utf8Reader)

		if err != nil {
			panic(err)
		}
		printCityList(all)
	}
}

func printCityList(contents []byte) {
	provinceList := make([]string, 0)
	re := regexp.MustCompile(`<a href=></a>`)

}

func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)

	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
