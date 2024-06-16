package util

import (
	"bili-down-fav/src/conf"
	"time"

	"github.com/go-resty/resty/v2"
)

const (
	// userAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.64 Safari/537.36"
	userAgent = ""
)

// 调用GetLoginInfo后，对全局变量cookie赋值
var CookieList = make(map[string]string)

func CookieStr() string {
	var str string
	for k, v := range CookieList {
		str += k + "=" + v + ";"
	}
	return str
}

// 全局客户端对象
var Http *resty.Client

func init() {
	CookieList = conf.List("cookie")
	Http = resty.New()
	Http.SetTimeout(time.Hour)
	Http.SetHeader("user-agent", userAgent)
	Http.SetHeader("cookie", CookieStr())
}
