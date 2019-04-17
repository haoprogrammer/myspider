package engine

import (
	"haoprogrammer/myspider/crawler/fetcher"
	"log"
)

type SimpleEngine struct {
}

func (e SimpleEngine) Run(seeds ...Request) {
	// Engine 维护Request队列
	var requests []Request

	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		parserResult, err := worker(r)
		if err != nil {
			continue
		}

		// 添加parseResult所有的Request到requests
		requests = append(requests, parserResult.Requests...)

		for _, item := range parserResult.Items {

			log.Printf("Got Item %v", item)
		}
	}
}

func worker(r Request) (ParserResult, error) {
	log.Printf("fetching %s", r.Url)
	// Fetch每1个Request获得原始网页转换为UTF-8编码的原始文本数据
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher error"+"fetching url %s: %v", r.Url, err)
		return ParserResult{}, err
	}

	//parserResult := r.ParserFunc(body)
	// 解析原始网页文本数据
	return r.ParserFunc(body), nil
}
