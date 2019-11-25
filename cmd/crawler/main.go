package main

import (
	"github.com/VicdVud/deli-crawler/internal/global"
	"github.com/VicdVud/deli-crawler/internal/xlsx"
	_ "github.com/spf13/viper"
	"log"
)

func init() {
	global.Init()
}

func main() {
	log.Println("deli-crawler start...")

	//crawler.DoCrawler()
	err := xlsx.ReadAndSave("E:/GitHub/deli-crawler/excel/2019-11-20.xlsx")
	if err != nil {
		log.Println(err)
	}
}
