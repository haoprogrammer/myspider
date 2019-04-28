package main

import (
	"fmt"
	"haoprogrammer/myspider/crawler_distributed/config"
	"haoprogrammer/myspider/crawler_distributed/rpcsupport"
	"haoprogrammer/myspider/crawler_distributed/worker"
	"testing"
	"time"
)

func TestCrawlService(t *testing.T) {
	const host = ":9000"
	go rpcsupport.ServeRPC(
		host, worker.CrawlService{})

	time.Sleep(time.Minute)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}
	req := worker.Request{
		Url: "http://album.zhenai.com/u/1518474013",
		Parser: worker.SerializedParser{
			Name: config.ParseProfile,
			Args: "阳儿",
		},
	}
	var result worker.ParseResult

	err = client.Call(config.CrawlServiceRpc, req, &result)

	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(result)
	}

}
