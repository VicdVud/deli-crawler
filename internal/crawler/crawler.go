package crawler

import (
	"github.com/spf13/viper"
	"log"
)

// DoCrawler 开始爬取
func DoCrawler() {
	log.Println("start crawling...")

	// 穿创建工作池
	concurrencyNum := viper.GetInt("crawl.concurrency_num")
	workerPool := NewPool(concurrencyNum)

	worker := NewCollyCrawler()
	workerPool.Run(worker)

	// 等待工作结束
	workerPool.Shutdown()
}