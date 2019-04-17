package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	. "golang.org/x/text/transform"
	"haoprogrammer/myspider/crawler/engine"
	"haoprogrammer/myspider/crawler/scheduler"
	"haoprogrammer/myspider/crawler/zhenai/parser"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
)

func getMsg() {
	resp, err := http.Get("https://www.zhenai.com/zhenghun")
	if err != nil {
		panic("occured a error")
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code", resp.StatusCode)
		return
	}

	//根据网站编码
	encoding := determineEncoding(resp.Body)
	utf8Reader := NewReader(resp.Body, encoding.NewDecoder())
	all, err := ioutil.ReadAll(utf8Reader)

	if err != nil {
		panic("occured a error")
	}

	//fmt.Printf("%s\n",all)
	printCityList(all)

}

//判断字符编码
func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	encoding, _, _ := charset.DetermineEncoding(bytes, "")

	return encoding
}

//打印城市列表
func printCityList(contents []byte) {
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
	matches := re.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		//for _,subMatch := range m {
		//	//fmt.Printf("%s ", subMatch)
		//
		//}
		fmt.Printf("City: %s, URL: %s\n", m[2], m[1])

		//fmt.Printf("%s\n", m)
		fmt.Println()
	}

	fmt.Println("Matches found %d\n", len(matches))
}

func main() {
	//getMsg()
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 10,
	}
	e.Run(engine.Request{
		Url:        "https://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})

	//resp, err := http.Get("http://album.zhenai.com/u/1194708821")

}
