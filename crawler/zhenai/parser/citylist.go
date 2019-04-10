package parser

import (
	"haoprogrammer/myspider/crawler/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParserCityList(contents []byte) engine.ParserResult{

	re := regexp.MustCompile(cityListRe)
	matchs := re.FindAllSubmatch(contents, -1)

	result := engine.ParserResult{}

	for _, m := range matchs  {
		//这里指明item类型为string
		result.Items = append(result.Items, string(m[2]))

		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParserFunc: engine.NilParser,
		})

	}
    return result
}


