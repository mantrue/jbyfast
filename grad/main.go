package main

import (
	"fmt"
	"net/http"
)

type Request struct {
	// HTTP请求
	httpReq *http.Request
	//请求的深度
	depth uint32
}

//用于创建一个新的请求实例
func NewRequest(httpReq *http.Request, depth uint32) *Request {
	return &Request{httpReq: httpReq, depth: depth}
}

//用于获取 HTTP 请求
func (req *Request) HTTPReq() *http.Request {
	return req.httpReq
}

//用于获取请求的深度
func (req *Request) Depth() uint32 {
	return req.depth
}

//数据响应的类型
type Response struct {
	// HTTP响应
	httpResp *http.Response
	//响应的深度
	depth uint32
}

//用于创建一个新的响应实例
func NewResponse(httpResp *http.Response, depth uint32) *Response {
	return &Response{httpResp: httpResp, depth: depth}
}

//用于获取HTTP响应
func (resp *Response) HTTPResp() *http.Response {
	return resp.httpResp
}

//用于获取响应深度
func (resp *Response) Depth() uint32 {
	return resp.depth
}

func main() {
	arr := []int{1, 2}

	//对于数组切片
	var tmp = [3]int{1, 2, 3}
	arr = tmp[0:2]
	fmt.Println(arr)
}
