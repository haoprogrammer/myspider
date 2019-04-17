package parser

import (
	"haoprogrammer/myspider/crawler/engine"
	"haoprogrammer/myspider/crawler/model"
	"regexp"
)

var ageRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([^<]+)岁</div>`)
var educationRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([^<]+)科</div>`)
var incomeRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798="">([^<]+)千</div>`)

func ParseProfile(contents []byte, name string) engine.ParserResult {
	profile := model.Profile{}

	profile.Name = name
	profile.Age = extractString(contents, ageRe)

	profile.Education = extractString(contents, educationRe)
	profile.Income = extractString(contents, incomeRe)

	result := engine.ParserResult{
		Items: []interface{}{profile},
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
