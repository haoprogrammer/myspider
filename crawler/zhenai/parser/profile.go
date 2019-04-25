package parser

import (
	"haoprogrammer/myspider/crawler/engine"
	"haoprogrammer/myspider/crawler/model"
	"regexp"
)

var ageRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([^<]+)岁</div>`)
var educationRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([^<]+)科</div>`)
var incomeRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798="">([^<]+)千</div>`)

var idUrlRe = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)

func parseProfile(contents []byte, url string, name string) engine.ParserResult {
	profile := model.Profile{}

	profile.Name = name
	profile.Age = extractString(contents, ageRe)

	profile.Education = extractString(contents, educationRe)
	profile.Income = extractString(contents, incomeRe)

	result := engine.ParserResult{
		Items: []engine.Item{
			{
				Url:     url,
				Type:    "zhenai",
				Id:      extractString([]byte(url), idUrlRe),
				Payload: profile,
			},
		},
	}

	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)

	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}

//抽取出公共方法
//func ProfileParser(name string) engine.ParserFunc {
//	//解析文本
//	return func(c []byte, url string) engine.ParserResult {
//		return ParseProfile(c, url, name)
//	}
//
//}

type ProfileParser struct {
	userName string
}

func (p *ProfileParser) Parse(contents []byte, url string) engine.ParserResult {
	return parseProfile(contents, url, p.userName)
}

func (p *ProfileParser) Serialize() (name string, args interface{}) {
	return "ProfileParser", p.userName
}

func NewProfileParser(name string) *ProfileParser {
	return &ProfileParser{
		userName: name,
	}
}
