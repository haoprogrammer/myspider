package parser

import (
	"io/ioutil"
	"testing"
)

func TestParserCityList(t *testing.T) {

	//contents, err := fetcher.Fetch("https://www.zhenai.com/zhenghun")
	//防止网络故障，或者说测试环境不一定能连到外网
	contents, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}
	result := ParseCityList(contents, "")
	//fmt.Printf("%s\n", contents)
	const resultSize = 470
	if len(result.Requests) != resultSize {
		t.Errorf("result should hava %d "+
			"Requests; but had %d", resultSize, len(result.Requests))
	}

	if len(result.Items) != resultSize {
		t.Errorf("result should hava %d "+
			"Items; but had %d", resultSize, len(result.Items))
	}

	expectedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}

	//expectedCities := []string{
	//	"City 阿坝",
	//	"City 阿克苏",
	//	"City 阿拉善盟",
	//}

	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected url %d: %s; but  was %s", i, url, result.Requests[i].Url)
		}
	}

	//for i, city := range expectedCities {
	//	if result.Items[i].(string) != city {
	//		t.Errorf("expected city %d: %s; but  was %s", i, city, result.Items[i].(string))
	//	}
	//}

}
