package global

import (
	"fmt"
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
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}

		// 初始deli数据
		InitDeliInfo()
	})
}
