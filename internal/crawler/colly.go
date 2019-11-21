package crawler

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
	Login.UserLogin()

	// 第二步，准备文件下载路径
	Prepare.PrepareFile()

	// 第三部，下载文件
	Download.DownloadFile()
}
