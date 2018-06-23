
package main

import(
	"github/golang_learning/golang_learning/src/flintlovesam/logger_v1"
)

func main(){
	logger_v1.InitLogger("console","INFO", "/User/xupeng/aa.txt")
	logger_v1.DEBUG("hhh")
	logger_v1.WARN("sss")
}
