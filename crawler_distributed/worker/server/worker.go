package main

import (
	"flag"
	"fmt"
	"haoprogrammer/myspider/crawler_distributed/rpcsupport"
	"haoprogrammer/myspider/crawler_distributed/worker"
	"log"
)

var port = flag.Int("port", 0, "the port for me to listen on")

func main() {

	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	//log.fatal打印输出内容
	//退出应用程序
	//defer函数不会执行
	log.Fatal(rpcsupport.ServeRPC(fmt.Sprintf(":%d", *port),
		worker.CrawlService{}))

}
