package main

import (
	"fmt"
	"github.com/olivere/elastic"
	"haoprogrammer/myspider/crawler_distributed/config"
	"haoprogrammer/myspider/crawler_distributed/persist"
	"haoprogrammer/myspider/crawler_distributed/rpcsupport"
	"log"
)

func main() {
	log.Fatal(serveRpc(
		fmt.Sprintf(":%d", config.ItemSarverPort),
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
