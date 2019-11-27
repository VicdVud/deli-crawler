package crawler

import (
	"encoding/json"
	"github.com/VicdVud/deli-crawler/internal/logger"
	"github.com/gocolly/colly"
)

type owner struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data []struct {
		OwnerId    string `json:"owner_id"`
		Id         string `json:"id"`
		Name       string `json:"name"`
		Type       int    `json:"type"`
		Area       string `json:"area"`
		Address    string `json:"address"`
		Industry   string `json:"industry"`
		Size       string `json:"size"`
		UpdateTime string `json:"update_time"`
		CreateTime string `json:"create_time"`
		OrgExtProp struct {
			StructureEnabled string `json:"structure_enabled"`
			AddUserMode      string `json:"add_user_mode"`
		} `json:"org_ext_prop"`
	} `json:"data"`
}

var ownerDefault = &owner{}

// FetchOwner 提取Owner
func (o *owner) FetchOwner() error {
	logger.Info("Start fetch owner...")

	cc := colly.NewCollector()
	var err error

	cc.OnRequest(func(r *colly.Request) {
		// 设置请求头
		r.Headers.Set("client_id", "eplus_web")
		r.Headers.Set("Accept", "application/json, text/plain, */*")
		r.Headers.Set("Cache-Control", "no-cache")
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36")
		r.Headers.Set("X-Service-Id", "organization")
		r.Headers.Set("Sec-Fetch-Site", "same-site")
		r.Headers.Set("Sec-Fetch-Mode", "cors")
		r.Headers.Set("Accept-Encoding", "gzip, deflate, br")
		r.Headers.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")
		r.Headers.Set("Authorization", loginDefault.Data.Token)
	})

	cc.OnResponse(func(r *colly.Response) {
		if 200 == r.StatusCode {
			if err = json.Unmarshal(r.Body, o); err == nil {
				logger.Info("Fetch owner succeeded")
				return
			}
		}
		logger.Info("Fetch owner failed: " + err.Error())
	})

	ownerUrl := "https://v2-app.delicloud.com/api/v2.0/org/findOrgDetailByUserId?user_id=" + loginDefault.Data.UserId
	return cc.Visit(ownerUrl)
}
