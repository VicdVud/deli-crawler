package global

import (
	"fmt"
	"github.com/VicdVud/deli-crawler/internal/logger"
	"github.com/spf13/viper"
	"sync"
)

var once = new(sync.Once)

func Init() {
	once.Do(func() {
		viper.SetConfigName("config")
		viper.AddConfigPath("/etc/crawler/")
		viper.AddConfigPath("$HOME/.crawler")
		viper.AddConfigPath(App.RootDir + "/config")

		err := viper.ReadInConfig()
		if err != nil {
			logger.Fatal(fmt.Errorf("Fatal error config file: %s \n", err))
		}

		// 设置日志文件夹
		logger.SetDir(App.RootDir + viper.GetString("log.dir"))

		// 初始deli数据
		InitDeliInfo()
	})
}
