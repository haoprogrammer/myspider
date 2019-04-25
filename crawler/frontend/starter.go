package main

import (
	"haoprogrammer/myspider/crawler/frontend/controller"
	"net/http"
)

func main() {

	http.Handle("/search",
		controller.CreateSearchResultHandler(
			"crawler/frontend/view/template.html"))
	err := http.ListenAndServe(":8888", nil)

	//如果服务器都没启动
	if err != nil {
		panic(err)
	}
}
