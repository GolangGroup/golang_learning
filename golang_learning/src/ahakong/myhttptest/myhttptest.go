package main

import (
	"net/http"
	"html/template"
	"fmt"
)

func login(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	if method == "GET" {
		w.Header().Set("Content-Type","text/html;charset=utf-8" )
		t, err := template.ParseFiles("./login2.html")
		if err != nil {
			fmt.Fprintf(w, "load login.html failed")
			return
		}
		t.Execute(w, nil)
	} else if method == "POST" {
		r.ParseForm()
		username := r.FormValue("username")
		passwd := r.FormValue("passwd")
		fmt.Printf("username:%s\n", username)
		fmt.Printf("passwd:%s\n", passwd)	
		if username == "admin" && passwd == "admin123" {
			fmt.Fprintf(w, "user %s login success", username)
		} else {
			fmt.Fprintf(w, "user %s login failed", username)
		}
	}
}

func main() {
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("listen server failed,err:%v\n", err)
		return
	}
}

