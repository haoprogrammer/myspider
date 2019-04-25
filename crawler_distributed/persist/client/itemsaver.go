package client

import (
	"haoprogrammer/myspider/crawler/engine"
	"haoprogrammer/myspider/crawler_distributed/config"
	"haoprogrammer/myspider/crawler_distributed/rpcsupport"
	"log"
)

func ItemSaver(host string) (chan engine.Item, error) {
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		return nil, err
	}
	out := make(chan engine.Item)
	//写save逻辑
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver:got item #%d:%v", itemCount, item)
			itemCount++

			//收到信息后,rpc通知做存储
			//Call rpc to save item
			result := ""
			err := client.Call(
				config.ItemSaverRpc,
				item, &result)
			if err != nil {
				log.Printf("item save error "+
					"saveing item %v: %v", item, err)
			}
		}
	}()
	return out, nil
}
