package logger_v1

import (
	"fmt"
	"time"
	"path"
)

type LoggerConsole struct{
	level int
}

func InitConsoleLogger(level int) loggerInterface {
	return &LoggerConsole{
		level:level,
	}
}

func printlog(level int, format string, args...interface{}){
	str:=fmt.Sprintf(format, args...)
	now:=time.Now()
	filename, funcname,lineno := getlineinfo()
	filename=path.Base(filename)
	filename=path.Base(funcname)
	fmt.Printf("%d/%d/%d/%d:%d:%d,%s,%s,%s:%d [%s]\r\n",now.Year(),now.Month(),now.Day(),
		now.Hour(), now.Minute(), now.Second(), leveltolevelstr(level), filename, funcname, lineno, str)
}


func (c *LoggerConsole) DEBUG(format string, args...interface{}){
	if c.level > LOGLEVELDEBUG{
		return
	}
	printlog(LOGLEVELDEBUG, format, args...)
}

func (c *LoggerConsole) TRACE(format string, args...interface{}){
	if c.level > LOGLEVELTRACE{
		return
	}
	printlog(LOGLEVELDEBUG, format, args...)
}

func (c *LoggerConsole) INFO(format string, args...interface{}){
	if c.level > LOGLEVELINFO{
		return
	}
	printlog(LOGLEVELDEBUG, format, args...)
}

func (c *LoggerConsole) WARN(format string, args...interface{}){
	if c.level > LOGLEVELWARN{
		return
	}
	printlog(LOGLEVELDEBUG, format, args...)
}

func (c *LoggerConsole) ERROR(format string, args...interface{}){
	if c.level> LOGLEVELERROR{
		return
	}
	printlog(LOGLEVELDEBUG, format, args...)
}

func (c *LoggerConsole) FATAL(format string, args...interface{}){
	if c.level> LOGLEVELFATAL{
		return
	}
	printlog(LOGLEVELDEBUG, format, args...)
}

func (f *LoggerFile)CLOSE(){
	f.fd.Close()

}