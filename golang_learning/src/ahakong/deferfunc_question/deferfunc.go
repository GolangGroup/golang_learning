package main

import (
	"fmt"
)

func defer_func1(i int)(t int) {
	t = i
	defer func(){
		t += 3
	}()
	return t
}

func defer_func2(i int)int {
	t := i
	defer func(){
		t += 3
	}()
	return t
}

func defer_func3(i int)(t int) {
	defer func(){
		t += i
	}()
	return 2
}

func main() {
	fmt.Println(defer_func1(10))
	fmt.Println(defer_func2(20))
	fmt.Println(defer_func3(30))
}