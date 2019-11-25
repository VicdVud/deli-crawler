package crawler

import (
	"encoding/json"
	"fmt"
	"github.com/VicdVud/deli-crawler/internal/global"
	"github.com/gocolly/colly"
	"log"
	"strings"
)

type exportExcel struct {
	Errno  string `json:"errno"`
	Errmsg string `json:"errmsg"`
	Data   struct {
		Url string `json:"url"`
	} `json:"data"`
}

var exportExcelDefault = &exportExcel{}
var exportCookie string

// 导出日期
var exportDate struct {
	year  int
	month int
	day   int
}

// ExportExcelFile 导出excel文件路径
// @brief year 考勤年份
// @brief month 考勤月份
// @brief day 考勤日期
func (e *exportExcel) ExportExcelFile(year, month, day int) error {
	log.Println("Start export excel...")

	exportDate.year = year
	exportDate.month = month
	exportDate.day = day

	cc := colly.NewCollector()
	var err error
	exportCookie = "gr_user_id=" + global.Deli.Cookie.GrUserId +
		"; sensorsdata2015jssdkcross=" + global.Deli.Cookie.Sensorsdata +
		"; grwng_uid=" + global.Deli.Cookie.GrwngUid +
		"; eplus_token=" + loginDefault.Data.Token +
		"; eplus_uid=" + loginDefault.Data.UserId +
		"; eplus_orgid=" + ownerDefault.Data[0].Id +
		"; eplus_origin_member_id=" + loginDefault.Data.UserId +
		"; PHPSESSID=" + phpSessionDefault.PhpSessId +
		"; __eplus_uid__=" + loginDefault.Data.UserId +
		"; __eplus_org_id__=" + ownerDefault.Data[0].Id +
		"; __eplus_token__=" + loginDefault.Data.Token +
		"; gr_session_id_8926433e4893eb23=" + global.Deli.Cookie.GrSessionId_8 +
		"; 95c6240b34dbd8be_gr_session_id=" + global.Deli.Cookie.CGrSessionId +
		"; 95c6240b34dbd8be_gr_session_id_53b4ddae-d8f1-4bee-a791-7bb54a86177d=" + global.Deli.Cookie.CGrSessionId53 +
		"; gr_session_id_95c6240b34dbd8be=" + global.Deli.Cookie.GrSessionId_9 +
		"; gr_session_id_95c6240b34dbd8be_a68070f4-223e-4fdd-8e96-39647fd8f564=" + global.Deli.Cookie.GrSessionId_95

	cc.OnRequest(func(r *colly.Request) {
		// 设置请求头
		r.Headers.Set("Accept", "application/json, text/javascript, */*; q=0.01")
		r.Headers.Set("X-Requested-With", "XMLHttpRequest")
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.97 Safari/537.36")
		r.Headers.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
		r.Headers.Set("Sec-Fetch-Site", "same-origin")
		r.Headers.Set("Sec-Fetch-Mode", "cors")
		r.Headers.Set("Referer", "https://v2-kq.delicloud.com/attend/admin/check/day")
		r.Headers.Set("Accept-Encoding", "gzip, deflate, br")
		r.Headers.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")

		// 设置Cookie
		r.Headers.Set("Cookie", exportCookie)
	})

	cc.OnResponse(func(r *colly.Response) {
		if 200 == r.StatusCode {
			if err = json.Unmarshal(r.Body, e); err == nil {
				// 此处有两种结果，需检查
				checkExportUrl()
				log.Println("Export excel succeeded")
				return
			}
		}
		log.Println("Export excel failed")
	})

	exportUrl := "https://v2-kq.delicloud.com/attend/admin/check/dayexport"
	exportBody := fmt.Sprintf("we_id=&tag_id=&staff_ids=&daterange=%d-%d-%d+-+%d-%d-%d&keyword=",
		year, month, day, year, month, day)

	return cc.PostRaw(exportUrl, []byte(exportBody))
}

// checkExportUrl 检查导出链接
// 有两种情况，当链接后缀为".xls"时，表明未导出过，为下载链接，可直接下载
// 否则重定向到导出列表，需重新导出后下载
func checkExportUrl() {
	// 若当前为文件下载链接，则后缀为".xls"
	if strings.Contains(exportExcelDefault.Data.Url, ".xls") {
		return
	}

	// 否则，则需重定向，再获取文件下载链接
	cc := colly.NewCollector()

	cc.OnRequest(func(r *colly.Request) {
		// 设置请求头
		r.Headers.Set("Upgrade-Insecure-Requests", "1")
		r.Headers.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3")
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36")
		r.Headers.Set("Sec-Fetch-User", "?1")
		r.Headers.Set("Sec-Fetch-Site", "same-origin")
		r.Headers.Set("Sec-Fetch-Mode", "nested-navigate")
		r.Headers.Set("Referer", "https://v2-kq.delicloud.com/attend/admin/check/day")
		r.Headers.Set("Accept-Encoding", "gzip, deflate, br")
		r.Headers.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")

		// 设置Cookie
		r.Headers.Set("Cookie", exportCookie)
	})

	cc.OnHTML("a[href]", func(e *colly.HTMLElement) {
		url := e.Attr("href")
		keyWord := fmt.Sprintf("%d年%d月%d日至%d年%d月%d日",
			exportDate.year, exportDate.month, exportDate.day,
			exportDate.year, exportDate.month, exportDate.day)
		if strings.Contains(url, keyWord) {
			// 重新导出excel报表
			// 获取导出id
			index := strings.Index(exportExcelDefault.Data.Url, "id=")
			if index != -1 {
				reexport(exportExcelDefault.Data.Url[index+3:])
			}

			exportExcelDefault.Data.Url = url
		}
	})

	url := "https://v2-kq.delicloud.com/" + exportExcelDefault.Data.Url
	cc.Visit(url)
}

// reexport 重新导出
// @param exportId 导出id
func reexport(exportId string) {
	cc := colly.NewCollector()

	cc.OnRequest(func(r *colly.Request) {
		// 设置请求头
		r.Headers.Set("Accept", "*/*")
		r.Headers.Set("X-Requested-With", "XMLHttpRequest")
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36")
		r.Headers.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
		r.Headers.Set("Sec-Fetch-Site", "same-origin")
		r.Headers.Set("Sec-Fetch-Mode", "cors")
		r.Headers.Set("Referer", "https://v2-kq.delicloud.com/admin/export/index")
		r.Headers.Set("Accept-Encoding", "gzip, deflate, br")
		r.Headers.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")

		// 设置Cookie
		r.Headers.Set("Cookie", exportCookie)
	})

	url := "https://v2-kq.delicloud.com/" + exportExcelDefault.Data.Url
	body := fmt.Sprintf("id=%s", exportId)

	cc.PostRaw(url, []byte(body))
	fmt.Println("abc")
}
