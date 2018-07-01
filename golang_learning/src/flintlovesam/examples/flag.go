package main

import (
	"flag"
	"fmt"

)

func main(){
	username := flag.String("name","", "Input your name")


	var flagint int
	flag.IntVar(&flagint, "flagname",12345 , "help name for flagname")

	flag.Parse()


	fmt.Println("Hello, ", *username)
	fmt.Println("ID = ", flagint)
}
