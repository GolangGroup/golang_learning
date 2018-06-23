package main

import (
	"ahakong/mylogger/logger"
	"time"
)

func main() {
	logger.InitLogger("file", "DEBUG", "C:\\github\\golang_learning\\golang_learning\\logger.txt")
	logger.DEBUG("this is a debug test")
	logger.WARN("this is a warn test")
	time.Sleep(time.Second)
}


