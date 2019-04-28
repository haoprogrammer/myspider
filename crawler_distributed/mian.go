package main

import (
	"bufio"
	"flag"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	. "golang.org/x/text/transform"
	"haoprogrammer/myspider/crawler/engine"
	"haoprogrammer/myspider/crawler/scheduler"
	"haoprogrammer/myspider/crawler/zhenai/parser"
	"haoprogrammer/myspider/crawler_distributed/config"
	itemsaver "haoprogrammer/myspider/crawler_distributed/persist/client"
	"haoprogrammer/myspider/crawler_distributed/rpcsupport"
	worker "haoprogrammer/myspider/crawler_distributed/worker/client"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/rpc"
	"regexp"
	"strings"
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

var (
	itemSaverHost = flag.String("itemsaver_host", "", "itemsaver host")

	workerHosts = flag.String("worker_hosts", "", "worker hosts (comma separated)")
)

func main() {
	//getMsg()
	//保证es启动，有存储数据的地方
	//itemChan, err := persist.ItemSaver("dating_profile")
	itemChan, err := itemsaver.ItemSaver(*itemSaverHost)
	if err != nil {
		panic(err)
	}

	pool := createClientPool(strings.Split(*workerHosts, ","))

	processor := worker.CreateProcessor(pool)

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}
	e.Run(engine.Request{
		Url: "https://www.zhenai.com/zhenghun",
		//ParserFunc: parser.ParseCityList,
		Parser: engine.NewFuncParser(
			parser.ParseCityList, config.ParseCityList),
	})
	//e.Run(engine.Request{
	//	Url:        "https://www.zhenai.com/zhenghun/shanghai",
	//	ParserFunc: parser.ParseCity,
	//})
	//
	//resp, err := http.Get("http://album.zhenai.com/u/1194708821")
	//getUserAgent()
}

func createClientPool(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client

	for _, h := range hosts {
		client, err := rpcsupport.NewClient(h)
		if err == nil {
			clients = append(clients, client)
			log.Printf("Connected to %s", h)
		} else {
			log.Printf("Error connecting to %s: %v", h, err)
		}

	}
	//往chan 分发
	out := make(chan *rpc.Client)
	//分发在goroutine里面，常用的套路写法
	go func() {

		//再套一层for，防止一轮分发完毕后不再分发
		for {
			//轮流分发
			for _, client := range clients {
				out <- client
			}
		}
	}()

	return out
}
