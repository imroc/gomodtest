package gomodtest

import "github.com/imroc/req/v2"

func Get() {
	req.Get("https://baidu.com/")
}

func Get2() {
	req.Get("https://baidu.com/")
}
