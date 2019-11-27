package main

import (
	"github.com/VicdVud/deli-crawler/internal/crawler"
	"github.com/VicdVud/deli-crawler/internal/global"
	"github.com/VicdVud/deli-crawler/internal/logger"
	"github.com/VicdVud/deli-crawler/internal/utils"
	"github.com/robfig/cron"
	"github.com/spf13/viper"
)

func main() {
	logger.Info("deli-crawler start...")

	if global.CrawlRegularly() {
		// 定时爬取
		c := cron.New()
		viper.SetDefault("crawl.spec", "0 0 */1 * * *")
		c.AddFunc(viper.GetString("crawl.spec"), func() {
			crawler.DoCrawler()
		})
		c.Start()
		defer c.Stop()

		// 等待系统中断信号
		utils.WaitSystemInterrupt()
	} else {
		// 单次爬取
		crawler.DoCrawler()
	}
}
