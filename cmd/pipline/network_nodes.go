package pipline

import (
	"bufio"
	"net"
)

//网络版中NetworkSink不要工作，（比如做等待等操作）
func NetworktSink(addr string, in <-chan int) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	// 这里用go func 让goroutine在后面做任务
	go func() {
		defer listener.Close()
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		defer conn.Close()

		writer := bufio.NewWriter(conn)
		//记得flush
		defer writer.Flush()
		WriterSink(writer, in)

	}()

}

func NetworkSource(addr string) <-chan int {
	out := make(chan int)
	go func() {
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			panic(err)
		}
		r := ReaderSourceChunk(bufio.NewReader(conn), -1)
		for v := range r {
			out <- v
		}
		close(out)
	}()
	return out

}
