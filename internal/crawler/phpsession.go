package crawler

import (
	"errors"
	"github.com/VicdVud/deli-crawler/internal/global"
	"github.com/VicdVud/deli-crawler/internal/logger"
	"github.com/gocolly/colly"
	"strings"
)

type phpSession struct {
	PhpSessId string
}

var phpSessionDefault = &phpSession{}

func (p *phpSession) FetchPhpSession() error {
	logger.Info("Start fetch PHPSESSID...")

	cc := colly.NewCollector()
	var err error

	cc.OnRequest(func(r *colly.Request) {
		// 设置请求头
		r.Headers.Set("Upgrade-Insecure-Requests", "1")
		r.Headers.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3")
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36")
		r.Headers.Set("X-Service-Id", "organization")
		r.Headers.Set("Sec-Fetch-User", "?1")
		r.Headers.Set("Sec-Fetch-Site", "same-site")
		r.Headers.Set("Sec-Fetch-Mode", "navigate")
		r.Headers.Set("Accept-Encoding", "gzip, deflate, br")
		r.Headers.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")

		// 设置Cookie
		cookie := "gr_user_id=" + global.Deli.Cookie.GrUserId +
			"; sensorsdata2015jssdkcross=" + global.Deli.Cookie.Sensorsdata +
			"; grwng_uid=" + global.Deli.Cookie.GrwngUid +
			"; eplus_token=" + loginDefault.Data.Token +
			"; eplus_uid=" + loginDefault.Data.UserId +
			"; eplus_orgid=" + ownerDefault.Data[0].Id +
			"; eplus_origin_member_id=" + loginDefault.Data.UserId

		r.Headers.Set("Cookie", cookie)
	})

	cc.OnResponse(func(r *colly.Response) {
		if 200 == r.StatusCode {
			// 因为此时响应的结果是重定向后的结果
			// 故可从重定向的请求头Cookie里找到PHPSESSID
			p.PhpSessId = parserSessId(r.Request.Headers.Get("Cookie"))
			if p.PhpSessId != "" {
				logger.Info("Fetch PHPSESSID succeeded")
				return
			}
		}
		err = errors.New("Cannot fetch PHPSESSID")
		logger.Info("Fetch PHPSESSID failed")
	})

	if err != nil {
		return err
	}

	ownerUrl := "https://v2-kq.delicloud.com/delicloud/login?token=" + loginDefault.Data.Token +
		"&user_id=" + loginDefault.Data.UserId +
		"&org_id=" + ownerDefault.Data[0].Id +
		"&origin_member_id=" + loginDefault.Data.UserId
	return cc.Visit(ownerUrl)
}

func parserSessId(cookie string) string {
	// 查询"PHPSESSID="开头的字符串
	index := strings.Index(cookie, "PHPSESSID=")
	cookie = cookie[index+10:]

	// 查询";"
	index = strings.Index(cookie, ";")
	if index != -1 {
		cookie = cookie[:index]
	}
	return cookie
}
