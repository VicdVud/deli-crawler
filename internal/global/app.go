package global

import (
	"github.com/VicdVud/deli-crawler/internal/utils"
	"os"
	"path/filepath"
	"sync"
)

type app struct {
	Name    string // 程序名称
	RootDir string // 项目根目录
	locker  sync.Locker
}

var App = &app{}

func init() {
	App.Name = os.Args[0]
	App.RootDir = inferRootDir()
}

func inferRootDir() string {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	var infer func(d string) string
	infer = func(d string) string {
		if utils.Exist(d + "/config") {
			return d
		}

		return infer(filepath.Dir(d))
	}

	return infer(cwd)
}
