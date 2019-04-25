package main

import (
	"fmt"
	"haoprogrammer/myspider/crawler_distributed/config"
	"haoprogrammer/myspider/crawler_distributed/rpcsupport"
	"haoprogrammer/myspider/crawler_distributed/worker"
	"log"
)

func main() {

	//log.fatal打印输出内容
	//退出应用程序
	//defer函数不会执行
	log.Fatal(rpcsupport.ServeRPC(fmt.Sprintf(":%d", config.WorkerPort0),
		worker.CrawService{}))

}
