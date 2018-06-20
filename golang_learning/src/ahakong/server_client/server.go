package main

import (
	"fmt"
	"net"
	"bufio"
	"os"
)

func main() {
	listenner, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Listen err=", err)
		return
	}
	
	defer listenner.Close()
	
	for {
		conn, err := listenner.Accept()
		if err != nil {
			fmt.Println("listenner.Accept err=", err)
			return
		}
		defer conn.Close()
		go func(conn net.Conn){
			fmt.Printf("地址 %s 接入服务器\n",conn.RemoteAddr())
			buf := make([]byte, 1024)
			n, err := conn.Read(buf)
			if err != nil {
				fmt.Println("conn.Read err=", err)
				return
			}
			fmt.Printf("收到消息：%s\n", string(buf[:n]))
			fmt.Println("请输入需要发送的信息：")
			reader := bufio.NewReader(os.Stdin)
			str, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("reader.ReadString err=", err)
				return
			}
			
			fmt.Println("发送中......")
			_, err = conn.Write([]byte(str))
			if err != nil {
				fmt.Println("reader.ReadString err=", err)
				return
			}
			fmt.Println("发送完毕，等待回复......")
			
		}(conn)
		
		
	}



}
