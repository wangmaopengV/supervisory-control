package main

import (
	"fmt"
	"net"
)

func main() {
	//主动连接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:9090")
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	defer conn.Close()
	//发送数据  因为服务器接受的数据为切片，所以要转换类型
	//conn.Write([]byte{0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01})

	exit := make(chan bool)
	exit <- true
}
