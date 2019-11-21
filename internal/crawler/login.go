package crawler

import (
	"encoding/json"
	"fmt"
	"github.com/VicdVud/deli-crawler/internal/global"
	"github.com/gocolly/colly"
	"log"
)

type login struct {
	token  string
	userId string
}

type LoginResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Token  string `json:"token"`
		Expire string `json:"expire"`
		UserId string `json:"user_id"`
	} `json:"data"`
}

var Login = &login{}

func (l *login) UserLogin() {
	log.Println("Start login...")

	cc := colly.NewCollector()

	cc.OnRequest(func(r *colly.Request) {
		// 设置请求头
		r.Headers.Set("client_id", "eplus_web")
		r.Headers.Set("Cache-Control", "no-cache")
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.97 Safari/537.36")
		r.Headers.Set("X-Service-Id", "userauth")
		r.Headers.Set("Content-Type", "application/json;charset=UTF-8")
		r.Headers.Set("Sec-Fetch-Site", "same-site")
		r.Headers.Set("Sec-Fetch-Mode", "cors")
		r.Headers.Set("Accept-Encoding", "gzip, deflate, br")
		r.Headers.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")
	})

	cc.OnResponse(func(r *colly.Response) {
		if 200 == r.StatusCode {
			var response LoginResponse
			if err := json.Unmarshal(r.Body, &response); err != nil {
				log.Fatal(err)
			}
			log.Println("Login succeeded")
		} else {
			log.Println("Login failed")
		}
	})

	loginUrl := "https://v2-app.delicloud.com/api/v2.0/auth/loginMobile"
	loginBody := fmt.Sprintf(`{"mobile":"%s","password":"%s"}`,
		global.Deli.Mobile,
		global.Deli.Password)

	err := cc.PostRaw(loginUrl, []byte(loginBody))
	if err != nil {
		log.Fatal(err)
	}
}
