package main

import (
	"haoprogrammer/myspider/crawler/engine"
	"haoprogrammer/myspider/crawler/model"
	"haoprogrammer/myspider/crawler_distributed/rpcsupport"
	"testing"
	"time"
)

func TestItemSaver(t *testing.T) {

	const host = ":1234"
	//start ItemSaverServer
	go serveRpc("host", "dating_profile")
	time.Sleep(time.Second)

	//start ItemSaverClient
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(nil)
	}

	//call save

	item := engine.Item{
		Url:  "http://album.zhenai.com/u/1518474013",
		Type: "zhenai",
		Id:   "1518474013",
		Payload: model.Profile{
			//Age: "51",
			Name: "阳儿",
		},
	}

	result := ""

	err = client.Call("ItemSaverService.Save", item, &result)

	if err != nil || result != "ok" {
		t.Errorf("result: %s , err: %s", result, err)
	}

}
