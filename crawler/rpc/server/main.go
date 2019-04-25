package main

import (
	"haoprogrammer/myspider/crawler/rpc"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	rpc.Register(rpcdemo.DemoService{})

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept error %v", err)
			continue
		}

		//在后台处理逻辑
		go jsonrpc.ServeConn(conn)
	}

}
