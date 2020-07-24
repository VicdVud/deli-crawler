package crawler

import (
	"fmt"
	"github.com/VicdVud/deli-crawler/internal/global"
	"github.com/VicdVud/deli-crawler/internal/logger"
	"github.com/VicdVud/deli-crawler/internal/model"
	"github.com/VicdVud/deli-crawler/internal/xlsx"
	"os"
	"strings"
	"time"
)

type CollyCrawler struct {
	timeDelay time.Duration // 每次抓取完成后停留时间，避免抓取太快导出被封
}

func NewCollyCrawler(d time.Duration) *CollyCrawler {
	return &CollyCrawler{
		timeDelay: d,
	}
}

func (c *CollyCrawler) Work() {
	// 第一步，登录账号
	err := loginDefault.UserLogin()
	if err != nil {
		logger.Fatal(err)
	}

	// 第二步，获取Owner信息
	err = ownerDefault.FetchOwner()
	if err != nil {
		logger.Fatal(err)
	}

	// 第三步，获取PHPSESSID
	err = phpSessionDefault.FetchPhpSession()
	if err != nil {
		logger.Fatal(err)
	}

	// 依据配置文件的导出起始时间与结束时间，执行不同的导出行为
	if global.Deli.FromDateString != global.TODAY {
		// 导出时间段记录
		for {
			// 判断日期是否截止
			dayEnd := global.Deli.FromDate.EqualTo(global.Deli.ToDate)

			// 导出下载
			c.ExportAndSave(global.Deli.FromDate)

			if dayEnd {
				// 日期截止，到处完后退出
				break
			}

			// 日期推后
			global.Deli.FromDate = global.Deli.FromDate.NextDay()

			// 强制延时
			time.Sleep(c.timeDelay)
		}
	} else {
		// 导出当天记录
		dayNow := time.Now()
		c.ExportAndSave(model.Date{
			Year:  dayNow.Year(),
			Month: int(dayNow.Month()),
			Day:   dayNow.Day(),
		})
	}
}

// ExportAndSave 导出excel文件并储存至数据库
func (c *CollyCrawler) ExportAndSave(date model.Date) {
	logger.Info("")
	logger.Info("+++++++++++++++++++++++")
	logger.Info("Exporting record: " + date.ToString())

	// 第四步，导出excel文件
	var err error

	// deli考勤报表导出机制有改变，在提交导出申请后，需等待一段时间后才可下载
	keyWord := fmt.Sprintf("%04d年%02d月%02d日至%04d年%02d月%02d日",
		date.Year, date.Month, date.Day,
		date.Year, date.Month, date.Day)

	// 重新获取次数限制
	exportMaxCount := 10
	exportCurCount := 0

	logger.Info("Start export excel...")

	for {
		exportCurCount++

		err = exportExcelDefault.ExportExcelFile(date)
		if err != nil {
			logger.Fatal(err)
		}

		if strings.Contains(exportExcelDefault.Data.Url, keyWord) {
			// 获取到导出的excel下载路径
			logger.Info("Export excel succeeded")
			break
		}

		if exportCurCount > exportMaxCount {
			// 导出次数达到上限，退出程序，汇报错误
			logger.Error("Cannot export excel: " + global.Deli.FromDate.ToString())
			os.Exit(0)
		}

		// 10S后重新获取
		logger.Info("Wait to export...")
		time.Sleep(10 * time.Second)
	}

	// 第五步，下载文件到本地
	var filePath string
	filePath, err = downloadDefault.DownloadFile(date)
	if err != nil {
		logger.Fatal(err)
	}

	// 第六步，储存至数据库
	err = xlsx.ReadAndSave(filePath)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Info("Save to db successfully, waiting next task...")
}
