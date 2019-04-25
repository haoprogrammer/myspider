package parser

import (
	"haoprogrammer/myspider/crawler/engine"
	"regexp"
)

//const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*><img src="[\s\S]*" alt="([^<]+)"></a>`
//const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

var (
	profileRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
)

//<a href="http://album.zhenai.com/u/1749079696" target="_blank"><img src="https://photo.zastatic.com/images/photo/437270/1749079696/20487218248971595.jpg?scrop=1&amp;crop=1&amp;w=140&amp;h=140&amp;cpos=north" alt="角落的泪光"></a>

func ParseCity(contents []byte, _ string) engine.ParserResult {

	//re := regexp.MustCompile(cityRe)
	matches := profileRe.FindAllSubmatch(contents, -1)

	result := engine.ParserResult{}

	for _, m := range matches {
		//url := string(m[1])
		////要把name拷贝出来
		//name := string(m[2])
		//这里指明item类型为string
		//2019417 可以不生成user的名字
		//result.Items = append(result.Items, "User "+name)

		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			//这里注意函数的作用域,函数执行,用原先的string(m[2])的值已经改变了
			//ParserFunc: ProfileParser(string(m[2])),
			Parser: NewProfileParser(string(m[2])),
		})

	}

	matches = cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests,
			engine.Request{
				Url: string(m[1]),
				//ParserFunc: ParseCity,
				Parser: engine.NewFuncParser(ParseCity, "ParseCity"),
			})
	}

	return result
}
