package parser

import (
	"haoprogrammer/myspider/crawler/engine"
	"haoprogrammer/myspider/crawler/model"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {

	contents, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}

	result := parseProfile(contents, "http://album.zhenai.com/u/1518474013", "阳儿")

	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 element;but was %v", result.Items)
	}

	actual := result.Items[0]

	expected := engine.Item{
		Url:  "http://album.zhenai.com/u/1518474013",
		Type: "zhenai",
		Id:   "1518474013",
		Payload: model.Profile{
			//Age: "51",
			Name: "阳儿",
		},
	}

	if actual != expected {
		t.Errorf("expected %v, but was %v", expected, actual)
	}
}
