package controller

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"haoprogrammer/myspider/crawler/engine"
	"haoprogrammer/myspider/crawler/frontend/model"
	"haoprogrammer/myspider/crawler/frontend/view"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

type SearchResultHandler struct {
	view view.SearchResultView

	client *elastic.Client
}

func CreateSearchResultHandler(template string) SearchResultHandler {

	client, err := elastic.NewClient(elastic.SetURL("http://10.10.55.113:31200"), elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	return SearchResultHandler{
		view:   view.CreateSearchResultView(template),
		client: client,
	}

}

//http:localhost:8888/search?q=男 已购房&from=20
func (h SearchResultHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	q := strings.TrimSpace(req.FormValue("q"))
	from, err := strconv.Atoi(req.FormValue("from"))

	if err != nil {
		from = 0
	}

	fmt.Fprintf(w, "q = %s,from %d", q, from)

	page, err := h.getSearchResult(q, from)
	fmt.Fprintf(w, "page = %v,from %d", page, from)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	err = h.view.Render(w, page)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

	}

}

func (h SearchResultHandler) getSearchResult(q string,
	from int) (model.SearchResult, error) {
	var result model.SearchResult
	resp, err := h.client.Search("dating_profile").
		Query(elastic.NewQueryStringQuery(q)).
		From(from).
		Do(context.Background())

	if err != nil {
		return result, err
	}

	result.Hits = resp.TotalHits()
	result.Start = from
	//go语言反射
	result.Items = resp.Each(
		reflect.TypeOf(engine.Item{}))

	return result, nil
}
