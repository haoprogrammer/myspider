package persist

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic"
	"haoprogrammer/myspider/crawler/engine"
	"haoprogrammer/myspider/crawler/model"
	"testing"
)

func TestItemSaver(t *testing.T) {

	expected := engine.Item{
		Url:  "http://album.zhenai.com/u/1518474013",
		Type: "zhenai",
		Id:   "1518474013",
		Payload: model.Profile{
			Age:  "51",
			Name: "阳儿",
		},
	}

	//TODO try to start up elastic search
	//here using docker go client
	client, err := elastic.NewClient(elastic.SetSniff(false),
		elastic.SetURL("http://10.10.55.113:31200"))
	//存储item
	const index = "dating_test"
	err = save(client, index, expected)
	if err != nil {
		panic(err)
	}

	//获取item
	resp, err := client.Get().
		Index(index).
		Type(expected.Type).
		Id(expected.Id).Do(context.Background())

	if err != nil {
		panic(err)
	}

	t.Logf("%s", resp.Source)

	var actual engine.Item
	err = json.Unmarshal([]byte(resp.Source), &actual)
	if err != nil {
		panic(err)
	}

	actualProfile, err := model.FormJsonObj(actual.Payload)
	actual.Payload = actualProfile

	if actual != expected {
		t.Errorf("got %v; expected %v",
			actual, expected)
	}
}
