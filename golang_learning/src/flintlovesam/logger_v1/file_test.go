package logger_v1

import "testing"

func TestInitFilelogger(t *testing.T) {
	InitLogger("console", "DEBUG", "/User/xupeng/logger.txt")
	DEBUG("fffff")
	WARN("flllll")
}
