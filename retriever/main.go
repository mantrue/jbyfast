package main

import (
	"fmt"
	real2 "retriever/real"
)

type Retiever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

func download(r Retiever) string {
	return r.Get("http://www.imooc.com")
}

func post(poster Poster) {
	poster.Post("http://www.imooc.com", map[string]string{"name": "geeker"})
}

type RetrieverPoster interface {
	Retiever
	Poster
}

const url = "http://www.imooc.com"

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{"contents": "another faend immcokc"})
	return s.Get("http://www.imooc.com")
}

func main() {
	var r Retiever
	r = real2.Retriever{}
	fmt.Println(download(r))
}
