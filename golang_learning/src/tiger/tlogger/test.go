package main

import(
	s "teamforgo/golang_learning/golang_learning/src/tiger/tlogger/version1"
	"time"
)

func main(){
	st := time.Now().Weekday().String();
	st = st +".log";
	s.InitLogger("file","INFO" , st);
	s.WARN("this is a test");
	s.TRACE("this is a test");
	s.DEBUG("this is a test");
	s.ERROR("this is a test");
	s.FATAL("this is a test");
	s.INFO("this is a test");
	s.CLOSE();
}