package httprun

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func SendPost(seturl string, parameter []byte) string {
	t := time.Now().Format("2006-01-02 15:04:05")
	token := md5.Sum([]byte("O1" + "O1" + t)) //算法加密
	tomd5 := fmt.Sprintf("%x", token)

	v := url.Values{}
	v.Set("appId", "O1")
	v.Add("time", t)
	v.Add("token", tomd5)
	//请求服务器url一定要记得进行编码操作

	//application/x-www-form-urlencoded
	resp, err := http.Post(seturl+v.Encode(),
		"application/json;charset=utf-8",
		bytes.NewBuffer(parameter))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return string(body)
}

func SendFormUnicode(seturl string, parameter []byte) string {
	t := time.Now().Format("2006-01-02 15:04:05")
	token := md5.Sum([]byte("O1" + "O1" + t)) //算法加密
	tomd5 := fmt.Sprintf("%x", token)

	v := url.Values{}
	v.Set("appId", "O1")
	v.Add("time", t)
	v.Add("token", tomd5)

	//fmt.Println(seturl + v.Encode())
	req, err := http.NewRequest("POST", seturl+v.Encode(), strings.NewReader(string(parameter)))
	if err != nil {
		panic(err)
	}

	// 表单方式(必须)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	//AJAX 方式请求
	req.Header.Add("x-requested-with", "XMLHttpRequest")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	return string(body)
}

func SendForm(seturl string, parameter []byte) string {
	t := time.Now().Format("2006-01-02 15:04:05")
	token := md5.Sum([]byte("O1" + "O1" + t)) //算法加密
	tomd5 := fmt.Sprintf("%x", token)

	v := url.Values{}
	v.Set("appId", "O1")
	v.Add("time", t)
	v.Add("token", tomd5)

	postValue := url.Values{
		"data": {string(parameter)},
	}

	postString := postValue.Encode()

	req, err := http.NewRequest("POST", seturl+v.Encode(), strings.NewReader(postString))
	if err != nil {
		// handle error
	}

	// 表单方式(必须)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	//AJAX 方式请求
	req.Header.Add("x-requested-with", "XMLHttpRequest")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	return string(body)
}
