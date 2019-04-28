package main

import (
	"flag"
	"fmt"
	"github.com/olivere/elastic"
	"haoprogrammer/myspider/crawler_distributed/config"
	"haoprogrammer/myspider/crawler_distributed/persist"
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

	log.Fatal(serveRpc(
		fmt.Sprintf(":%d", *port),
		config.ElasticIndex))
}

func serveRpc(host, index string) error {

	client, err := elastic.NewClient(elastic.SetSniff(false),
		elastic.SetURL("http://10.10.55.113:31200"))
	//存储item
	if err != nil {
		panic(err)
	}

	return rpcsupport.ServeRPC(host, &persist.ItemSaverService{
		Client: client,
		Index:  index,
	})
}
