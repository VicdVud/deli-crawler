package crawler

import (
	"fmt"
	"github.com/VicdVud/deli-crawler/internal/global"
	"github.com/VicdVud/deli-crawler/internal/model"
	"github.com/gocolly/colly"
	"log"
	"os"
)

type download struct {
}

var downloadDefault = &download{}

// DownloadFile 下载文件
// 返回文件保存路径
func (d *download) DownloadFile(date model.Date) (string, error) {
	log.Println("Download excel...")

	// 文件保存路径
	excelPath := global.Deli.SaveDir + fmt.Sprintf("%d-%d-%d.xlsx", date.Year, date.Month, date.Day)

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
			excel, err := os.OpenFile(excelPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
			if err != nil {
				log.Fatal("Cannot create file: " + excelPath)
			}
			defer excel.Close()

			_, err = excel.Write(r.Body)
			if err != nil {
				log.Fatal("Cannot write data to file: " + excelPath + "\n" + err.Error())
			}

			log.Println("Download excel succeeded")
		}
	})

	downloadUrl := "https://v2-kq.delicloud.com" + exportExcelDefault.Data.Url
	err := cc.Visit(downloadUrl)
	if err != nil {
		fmt.Println(err)
	}
	return excelPath, err
}
