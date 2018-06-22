package logger

import (
	"testing"
)

func TestFileLogger(t *testing.T) {
	InitLogger("console", LOGLEVELDEBUG, "C:\\github\\golang_learning\\golang_learning\\logger.txt")
	DEBUG("ffsf")
	WARN("fllll")

}

