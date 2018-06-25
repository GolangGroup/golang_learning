package main

import (
	"encoding/xml"
	"io/ioutil"
	"fmt"
)

type Servers struct {
	Name xml.Name `xml:"servers"`
	Version string `xml:"version,attr"`
	Servers []Server `xml:"server"`
}

type Server struct {
	ServerName string `xml:"serverName"`
	ServerIP string `xml:"serverIP"`
}

func main() {
	data, err := ioutil.ReadFile("./config.xml")
	if err != nil {
		fmt.Printf("read configxml failed err %v\n",err)
		return
	}
	fmt.Printf("%s\n", string(data[:]))

	var servers Servers
	err = xml.Unmarshal(data, &servers)
	if err != nil {
		fmt.Printf("Unmarshal failed err %v\n",err)
		return	
	}
	fmt.Printf("xml:%#v\n", servers)
}

