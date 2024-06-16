package user

import (
	"bili-down-fav/src/conf"
	"bili-down-fav/src/util"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/skip2/go-qrcode"
)

func Login(callback func()) {
	code := qr()
	check(code, callback)

}

func CurrentUser() *Info {
	url := "https://api.bilibili.com/x/web-interface/nav"
	userInfo := new(Info)
	if resp, err := util.Http.R().
		Get(url); err == nil {
		json.Unmarshal(resp.Body(), userInfo)
	}
	return userInfo
}

func check(qrcode string, callback func()) {
	url := "https://passport.bilibili.com/x/passport-login/web/qrcode/poll"
	loginInfo := new(LoginResp)
	var header http.Header
	for {
		if resp, err := util.Http.R().
			SetQueryParam("qrcode_key", qrcode).
			Get(url); err == nil {
			json.Unmarshal(resp.Body(), loginInfo)
			header = resp.Header()
		}
		if loginInfo.Data.Code == 0 {
			log.Println("登录成功")
			// 登录后会自动设置cookie到http-cli中，不需要手动放
			for _, v := range header.Values("set-cookie") {
				pair := strings.Split(v, ";")
				kv := strings.Split(pair[0], "=")
				util.CookieList[kv[0]] = kv[1]
				conf.Save("cookie", kv[0], kv[1])
			}
			callback()
			break
		} else {
			log.Println(loginInfo.Data.Message)
			if loginInfo.Data.Code == 86038 {
				break
			}
		}
		time.Sleep(time.Second)
	}
}

func qr() string {
	util.Http.Header.Del("cookie")
	url := "https://passport.bilibili.com/x/passport-login/web/qrcode/generate"
	qrUrl := new(QrUrl)
	if resp, err := util.Http.R().Get(url); err == nil {
		json.Unmarshal(resp.Body(), qrUrl)
	}
	if err := qrcode.WriteFile(qrUrl.Data.Url, qrcode.Medium, 256, conf.QrPath); err != nil {
		log.Println("生成二维码失败：", err)
		return ""
	}
	log.Println("二维码已生成至" + conf.QrPath)
	return qrUrl.Data.QrcodeKey
}
