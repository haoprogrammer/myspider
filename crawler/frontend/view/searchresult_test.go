package view

import (
	"haoprogrammer/myspider/crawler/engine"
	"haoprogrammer/myspider/crawler/frontend/model"
	common "haoprogrammer/myspider/crawler/model"
	"os"
	"testing"
)

func TestSearchResultView_Render(t *testing.T) {

	//template := template.Must(template.ParseFiles("template.html"))
	view := CreateSearchResultView("template.html")
	out, err := os.Create("template.test.html")
	page := model.SearchResult{}
	page.Hits = 123

	item := engine.Item{
		Url:  "http://album.zhenai.com/u/1518474013",
		Type: "zhenai",
		Id:   "1518474013",
		Payload: common.Profile{
			Age:  "51",
			Name: "阳儿",
		},
	}

	for i := 0; i < 10; i++ {
		page.Items = append(page.Items, item)
	}

	//err = template.Execute(out, page)
	err = view.Render(out, page)
	if err != nil {
		panic(err)
	}

}
