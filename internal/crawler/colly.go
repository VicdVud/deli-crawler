package crawler

import "log"

type CollyCrawler struct {
	waiter chan int
}

func NewCollyCrawler() *CollyCrawler {
	return &CollyCrawler{
		waiter: make(chan int),
	}
}

func (c *CollyCrawler) Work() {
	// 第一步，登录账号
	err := loginDefault.UserLogin()
	if err != nil {
		log.Fatal(err)
	}

	// 第二步，获取Owner信息
	err = ownerDefault.FetchOwner()
	if err != nil {
		log.Fatal(err)
	}

	// 第三步，获取PHPSESSID
	err = phpSessionDefault.FetchPhpSession()
	if err != nil {
		log.Fatal(err)
	}

	// 第四步，导出excel文件
	err = exportExcelDefault.ExportExcelFile(2019, 11, 20)
	if err != nil {
		log.Fatal(err)
	}

	// 第五步，下载文件到本地
	err = downloadDefault.DownloadFile(2019, 11, 20)
	if err != nil {
		log.Fatal(err)
	}
}
