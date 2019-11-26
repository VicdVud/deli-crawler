package global

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
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
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}

		// 设定日志文件
		logPath := App.RootDir + viper.GetString("log.path")
		logFile, err := os.OpenFile(logPath, os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			log.Fatal("Cannot open log file")
		}
		log.SetOutput(logFile)

		// 初始deli数据
		InitDeliInfo()
	})
}
