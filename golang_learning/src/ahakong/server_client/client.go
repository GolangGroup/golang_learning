package main

import (
	"fmt"
	"net"
	"bufio"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	
	defer conn.Close()
	buf := make([]byte, 1024)
	
	fmt.Printf("Usage:input the string to send, press `q` to quit\n")
	for {
		fmt.Println("请输入需要发送的信息：")
		reader := bufio.NewReader(os.Stdin)
		str, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("reader.ReadString err =", err)
			return
		}
		
		if str == "q\r\n" {
			fmt.Println("退出客户端程序")
			return
		}
		
		_, err = conn.Write([]byte(str))
		if err != nil {
			fmt.Println("conn.Write err=", err)
			return
		}
		fmt.Println("发送完毕，等待回复.....")
		
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("conn.Read err=", err)
			return
		}
		
		fmt.Printf("收到回复：%s\n", string(buf[:n]))
		
	}
	
}

