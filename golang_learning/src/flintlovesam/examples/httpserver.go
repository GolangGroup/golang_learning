package main

import (
	"net/http"
	"fmt"
	"strings"
	"log"
)

func sayhelloName(w http.ResponseWriter, r* http.Request){
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k ,v:=range r.Form{
		fmt.Println("key:",k)
		fmt.Println("val:",strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello xupeng come on")
}

func main(){
		http.HandleFunc("/" , sayhelloName)
		err := http.ListenAndServe(":7070",nil)

		if err != nil{
			log.Fatal("ListenAndServer", err)
		}

}