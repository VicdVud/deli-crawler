package crawler

import (
	"github.com/VicdVud/deli-crawler/internal/logger"
	"github.com/spf13/viper"
	"time"
)

// DoCrawler 开始爬取
func DoCrawler() {
	logger.Info("start crawling...")

	// 穿创建工作池
	concurrencyNum := viper.GetInt("crawl.concurrency_num")
	workerPool := NewPool(concurrencyNum)

	d, err := time.ParseDuration(viper.GetString("crawl.sleep"))
	if err != nil {
		// 解析错误，则默认停留10s
		d = 10 * time.Second
	}

	worker := NewCollyCrawler(d)
	workerPool.Run(worker)

	// 等待工作结束
	workerPool.Shutdown()
}
