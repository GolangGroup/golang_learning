package main

import(
	"fmt"
	"runtime"
)

func say(cont string){
	for i:= 0;i<10;i++ {
		runtime.Gosched()
		fmt.Println(cont+": the "+string(i)+" runs")
	}
}
func main(){
	runtime.GOMAXPROCS(4)
	go say("hello")
	say("world")
}