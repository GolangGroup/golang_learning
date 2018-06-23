package logger_v1

import (
	"fmt"
	"runtime"
)

const(
	LOGLEVELDEBUG=iota
	LOGLEVELTRACE
	LOGLEVELINFO
	LOGLEVELWARN
	LOGLEVELERROR
	LOGLEVELFATAL
)

func leveltolevelstr(level int)string{
	switch level {
	case LOGLEVELDEBUG:
		return "DEBUG"
	case LOGLEVELTRACE:
		return "TRACE"
	case LOGLEVELINFO:
		return "INFO"
	case LOGLEVELWARN:
		return "WARN"
	case LOGLEVELERROR:
		return "ERROR"
	case LOGLEVELFATAL:
		return "FATAL"

	default:
		fmt.Printf("leveltolevelstr err, level%d",level)
		return "DEBUG"

	}

}

func levelstrtolevel(levelstr string)int  {
	switch levelstr {
	case "DEBUG":
		return LOGLEVELDEBUG
	case "TRACE":
		return LOGLEVELTRACE
	case "INFO":
		return LOGLEVELINFO
	case "WARN":
		return LOGLEVELWARN
	case "ERROR":
		return LOGLEVELERROR
	case "FATAL":
		return LOGLEVELFATAL
	default: fmt.Printf("levelstrtolevel err,levelstr%s\n",levelstr)
		return LOGLEVELDEBUG
	}

}


var log loggerInterface

func InitLogger(logtype string, levelstr string, path string){
	level:=levelstrtolevel(levelstr)
	switch logtype {
	case "file":
		log = InitFilelogger(level, path)

	case "console":
		log=InitConsoleLogger(level)
	}
}

func DEBUG(format string, args...interface{}){
	log.DEBUG(format, args...)
}

func TRACE(format string, args...interface{}){
	log.TRACE(format, args...)
}

func INFO(format string, args...interface{}){
	log.INFO(format, args...)
}

func WARN(format string, args...interface{}){
	log.INFO(format, args...)
}

func ERROR(format string, args...interface{}){
	log.INFO(format, args...)
}

func FATAL(format string, args...interface{}){
	log.INFO(format, args...)
}

func getlineinfo()(filename string, funcname string, lineno int){

	pc,file,line,ok:=runtime.Caller(4)

	if ok{
		filename = file
		funcname = runtime.FuncForPC(pc).Name()
		lineno = line

	}
	return
}
