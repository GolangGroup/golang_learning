package logger

import (
	"fmt"
	"os"
	"time"
	"path"
)

type LoggerFile struct {
	level int
	fd *os.File
}

func InitFileLogger(level int, path string) LoggerInterface{
	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("InitFileLogger os.OpenFile err=", err)
		return nil
	}
	loggerFileObj := &LoggerFile{
		level: level,
		fd:file,
	}
	return loggerFileObj
}

func writelog(fd *os.File, level int, format string, args...interface{}) {
	str := fmt.Sprintf(format, args...)
	now := time.Now()
	filename, funcname, lineno := getlineinfo()
	filename = path.Base(filename)
	funcname = path.Base(funcname)
	fmt.Fprintf(fd, "%d/%d/%d/%d:%d:%d,%s,%s,%s:%d [%s]\r\n", now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second(), leveltolevelstr(level), filename, funcname, lineno, str)
}

func (f *LoggerFile) DEBUG(format string, args...interface{}) {
	if f.level > LOGLEVELDEBUG {
		return
	}
	writelog(f.fd, LOGLEVELDEBUG, format, args...)
}

func (f *LoggerFile) TRACE(format string, args...interface{}) {
	if f.level > LOGLEVELTRACE {
		return
	}	
	writelog(f.fd, LOGLEVELTRACE, format, args...)
}

func (f *LoggerFile) INFO(format string, args...interface{}) {
	if f.level > LOGLEVELINFO {
		return
	}	
	writelog(f.fd, LOGLEVELINFO, format, args...)
}

func (f *LoggerFile) WARN(format string, args...interface{}) {
	if f.level > LOGLEVELWARN {
		return
	}	
	writelog(f.fd, LOGLEVELWARN, format, args...)
}

func (f *LoggerFile) ERROR(format string, args...interface{}) {
	if f.level > LOGLEVELERROR {
		return
	}	
	writelog(f.fd, LOGLEVELERROR, format, args...)
}

func (f *LoggerFile) FATAL(format string, args...interface{}) {
	if f.level > LOGLEVELFATAL {
		return
	}	
	writelog(f.fd, LOGLEVELFATAL, format, args...)
}

func (f *LoggerFile) CLOSE() {
	f.fd.Close()
}