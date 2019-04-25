package persist

import (
	"context"
	"github.com/olivere/elastic"
	"github.com/pkg/errors"
	"haoprogrammer/myspider/crawler/engine"
	"log"
)

func ItemSaver(index string) (chan engine.Item, error) {
	client, err := elastic.NewClient(elastic.SetURL("http://10.10.55.113:31200"), elastic.SetSniff(false))
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

			//将数据存储到es中
			err := Save(client, index, item)
			if err != nil {
				log.Printf("item save error "+
					"saveing item %v: %v", item, err)
			}
		}
	}()
	return out, nil
}

//将item存到es中
func Save(client *elastic.Client, index string, item engine.Item) error {
	//判断type是否为空
	if item.Type == "" {
		return errors.New("must supply Type")
	}
	//index由配置人员配置
	//type程序给配置
	indexService := client.Index().
		Index(index).
		Type(item.Type).
		BodyJson(item)
	if item.Id != "" {
		indexService.Id(item.Id)
	}

	_, err := indexService.Do(context.Background())

	if err != nil {
		return nil
	}

	return nil
}
