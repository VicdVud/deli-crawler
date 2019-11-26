package utils

import (
	"log"
	"os"
	"os/signal"
)

// 等待系统中断
func WaitSystemInterrupt() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	<-c
	log.Println("System interrupted")
}
