package parser

import (
	"haoprogrammer/myspider/crawler/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

//所有的parser都传入url，不用url可以甩下划线表示
func ParseCityList(contents []byte, url string) engine.ParserResult {

	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParserResult{}
	limit := 10

	for _, m := range matches {
		//这里指明item类型为string
		//2019417可以不生成无价值的item
		//result.Items = append(result.Items, "City "+string(m[2]))

		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			//ParserFunc: engine.NilParser,
			//ParserFunc: ParseCity,
			Parser: engine.NewFuncParser(ParseCity, "ParseCity"),
		})
		limit--
		if limit <= 0 {
			break
		}
	}
	return result
}
