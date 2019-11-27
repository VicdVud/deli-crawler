package logger

import (
	"errors"
	"fmt"
	"github.com/VicdVud/deli-crawler/internal/model"
	"github.com/VicdVud/deli-crawler/internal/utils"
	"log"
	"os"
	"time"
)

type Logger struct {
	file    *FileLogger    // 文件流
	console *ConsoleLogger // 控制台流

	dir string // 日志文件目录
}

var loggerDefault = &Logger{
	dir:     "log/", // 默认日志文件夹
	file:    NewFileLogger(),
	console: NewConsoleLogger(),
}

// 自定义文件日志写入接口
type FileWriter struct {
	today   model.Date // 当前日志日期
	logFile *os.File
}

func (f *FileWriter) Write(p []byte) (n int, err error) {
	now := time.Now()
	today := model.Date{
		Year:  now.Year(),
		Month: int(now.Month()),
		Day:   now.Day(),
	}

	if f.logFile != nil && f.today.EqualTo(today) {
		// 已创建日志文件，且在同一天，则直接写入
		return f.logFile.Write(p)
	}

	// 避免文件夹未创建导致错误
	utils.CreateDir(loggerDefault.dir)

	// 创建日志文件并打印日志
	f.today = today
	f.logFile, err = generateLogFile(f.today)
	if err != nil {
		e := "Cannot create/open logfile on day: " + f.today.ToString()
		fmt.Println(e)
		return -1, errors.New(e)
	} else {
		return f.logFile.Write(p)
	}
}

// 控制台日志流
type ConsoleLogger struct {
	logFile *log.Logger
}

func NewConsoleLogger() *ConsoleLogger {
	return &ConsoleLogger{
		logFile: log.New(os.Stdout, "", log.Ldate|log.Ltime),
	}
}

// 文件日志流
type FileLogger struct {
	logFile *log.Logger
}

func NewFileLogger() *FileLogger {
	f := &FileLogger{logFile: nil}
	fw := &FileWriter{}
	f.logFile = log.New(fw, "", log.Ldate|log.Ltime)
	return f
}

// 依据日期生成日志文件，命名格式为"2019-01-01.log"
func generateLogFile(date model.Date) (*os.File, error) {
	logName := fmt.Sprintf("%04d-%02d-%02d.log", date.Year, date.Month, date.Day)
	return os.OpenFile(loggerDefault.dir+logName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
}

func Debug(v ...interface{}) {
	str := "[DEBUG] " + fmt.Sprint(v...)
	_ = loggerDefault.file.logFile.Output(4, str)
	_ = loggerDefault.console.logFile.Output(4, str)
}

func Info(v ...interface{}) {
	str := "[INFO] " + fmt.Sprint(v...)
	_ = loggerDefault.file.logFile.Output(4, str)
	_ = loggerDefault.console.logFile.Output(4, str)
}

func Warn(v ...interface{}) {
	str := "[WARN] " + fmt.Sprint(v...)
	_ = loggerDefault.file.logFile.Output(4, str)
	_ = loggerDefault.console.logFile.Output(4, str)
}

func Error(v ...interface{}) {
	str := "[ERROR] " + fmt.Sprint(v...)
	_ = loggerDefault.file.logFile.Output(4, str)
	_ = loggerDefault.console.logFile.Output(4, str)
}

func Fatal(v ...interface{}) {
	str := "[FATAL] " + fmt.Sprint(v...)
	_ = loggerDefault.file.logFile.Output(4, str)
	_ = loggerDefault.console.logFile.Output(4, str)
	os.Exit(1)
}

// 设置log文件夹
func SetDir(dir string) {
	loggerDefault.dir = dir
	utils.CreateDir(dir)
}
