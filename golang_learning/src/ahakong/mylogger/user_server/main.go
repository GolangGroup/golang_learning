package main

import (
	"ahakong/mylogger/logger"
	
)

func main() {
	logger.InitLogger("console", "INFO", "C:\\github\\golang_learning\\golang_learning\\logger.txt")
	logger.DEBUG("this is a debug test")
	logger.WARN("this is a warn test")
}


