package rpcsupport

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

//RPC服务端
func ServeRPC(host string, service interface{}) error {
	rpc.Register(service)

	listener, err := net.Listen("tcp", host)
	if err != nil {
		return err
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
	return nil
}

//jsonrpc客户端使用
func NewClient(host string) (*rpc.Client, error) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return nil, err
	}

	client := jsonrpc.NewClient(conn)

	return client, nil
}
