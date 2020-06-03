package main

import (
	"log"
	"net"
	"net/rpc"
)

func main() {
	// 将对象类型中所有满足PRC规则的对方方法注册为RPC函数
	err := rpc.RegisterName("HelloService", new(HelloService))
	if err != nil {
		log.Fatal("Register server error: ", err)
	}

	// 建立TCP连接
	listenner, err := net.Listen("tcp", ":1")
	if err != nil {
		log.Fatal("Listen TCP error: ", err)
	}

	for {
		// 监听连接
		conn, err := listenner.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		// 提供RPC服务
		go rpc.ServeConn(conn)
	}

}
