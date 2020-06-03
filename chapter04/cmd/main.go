package main

import "net/rpc"

func main() {
	rpc.RegisterName("HelloService", new(server.HelloService))
	
}
