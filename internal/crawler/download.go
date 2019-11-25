package crawler

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"os"
)

type download struct {
}

var downloadDefault = &download{}

func (d *download) DownloadFile(year, month, day int) error {
	log.Println("Download excel...")

	cc := colly.NewCollector()

	cc.OnRequest(func(r *colly.Request) {
		// 设置请求头
		r.Headers.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3")
		r.Headers.Set("X-Requested-With", "XMLHttpRequest")
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.97 Safari/537.36")
		r.Headers.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
		r.Headers.Set("Sec-Fetch-User", "?1")
		r.Headers.Set("Sec-Fetch-Site", "same-origin")
		r.Headers.Set("Sec-Fetch-Mode", "navigate")
		r.Headers.Set("Accept-Encoding", "gzip, deflate, br")
		r.Headers.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")

		// 设置Cookie
		r.Headers.Set("Cookie", exportCookie)
	})

	cc.OnResponse(func(r *colly.Response) {
		if 200 == r.StatusCode {
			excelName := fmt.Sprintf("%d-%d-%d.xlsx", year, month, day)
			excel, err := os.OpenFile(excelName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
			if err != nil {
				log.Fatal("Cannot create file: " + excelName)
			}
			defer excel.Close()

			_, err = excel.Write(r.Body)
			if err != nil {
				log.Fatal("Cannot write data to file: " + excelName)
			}
		}
	})

	downloadUrl := "https://v2-kq.delicloud.com" + exportExcelDefault.Data.Url
	err := cc.Visit(downloadUrl)
	if err != nil {
		fmt.Println(err)
	}
	return err
}
