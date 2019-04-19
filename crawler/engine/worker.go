package engine

import (
	"haoprogrammer/myspider/crawler/fetcher"
	"log"
)

//1.通过url拿内容
//2.调用parser函数，返回内容ParserResult
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
	return r.ParserFunc(body, r.Url), nil
}
