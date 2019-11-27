package utils

import "os"

// Exist 检查文件或目录是否存在
// 如果由 filename 指定的文件或目录存在则返回 true，否则返回 false
func Exist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

// 创建文件夹
func CreateDir(dir string) {
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		// 文件夹不存在, 则创建
		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			panic("Cannot create log dir, please check again")
		}
	}
}
