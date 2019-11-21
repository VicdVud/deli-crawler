package main

import (
	"github.com/VicdVud/deli-crawler/internal/crawler"
	"github.com/VicdVud/deli-crawler/internal/global"
	_ "github.com/spf13/viper"
	"log"
)

func init() {
	global.Init()
}

func main() {
	log.Println("deli-crawler start...")

	crawler.DoCrawler()
}
