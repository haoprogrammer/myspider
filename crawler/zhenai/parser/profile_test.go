package parser

import (
	"haoprogrammer/myspider/crawler/model"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {

	contents, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}

	result := ParseProfile(contents, "静听雨声")

	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 element;but was %v", result.Items)
	}

	profile := result.Items[0].(model.Profile)

	expected := model.Profile{
		Age: string(26),
	}

	if profile != expected {
		t.Errorf("expected %v, but was %v", expected, profile)
	}
}
