package main

import (
	"reflect"
	"fmt"
)

type User struct {
	Id int
	Name string
	Age int
}

func (u User) Hello(name string) string{
	fmt.Println("Hello",name, "my name is ",u.Name)
	return fmt.Sprintf("%p", &u)
}

type Manager struct {
	User
	title string

}

func main(){
	m := Manager{User:User{1, "OK", 12}, title:"123"}
	t := reflect.TypeOf(m)

	fmt.Printf("get anoymous field = %v\n", t.FieldByIndex([]int{0,0}))
	fmt.Printf("get anoymous field = %v\n", t.FieldByIndex([]int{0,1}))

	num := 2345
	v:=reflect.ValueOf(&num)
	v.Elem().SetInt(1234)

	fmt.Println(num)


	u:= User{1, "OK", 12}
	Set(&u)
	fmt.Println(u)
	vv:= reflect.ValueOf(u)
	mv:=vv.MethodByName("Hello")
	args:=[]reflect.Value{reflect.ValueOf("xupeng\n")}
	mv.Call(args)


	myType := &User{2,"xupeng",25}
	fmt.Println("####################")
	fmt.Println(myType)
	fmt.Println("####################")
	mtV := reflect.ValueOf(&myType).Elem()
	fmt.Println("####################")
	fmt.Println(mtV)
	fmt.Println("####################")

	myType2 := User{3,"xupeng",25}

	mtV2 := reflect.ValueOf(myType2)
	params2 := make([]reflect.Value, 1)
	params2[0] = reflect.ValueOf("XUpeng\n")
	//params2:=[]reflect.Value{reflect.ValueOf("xupeng\n")}
	//mtV2.MethodByName("hello").Call(params2)
	mtV2.MethodByName("Hello").Call(params2)
	

	fmt.Println("return=-----", mtV.MethodByName("Hello").Call(params2)[0])

}

func Set(o interface{}){

	v := reflect.ValueOf(o)
	if v.Kind()==reflect.Ptr && !v.Elem().CanSet(){
		fmt.Println("cant set")

	}else{
		v = v.Elem()

	}
	f:=v.FieldByName("Name")
	if !f.IsValid(){
		fmt.Println("BAD")
		return
	}

	if f:=v.FieldByName("Name"); f.Kind() == reflect.String{
		f.SetString("BYE")

	}
}