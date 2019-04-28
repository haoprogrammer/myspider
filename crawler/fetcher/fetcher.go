package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"haoprogrammer/myspider/crawler_distributed/config"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

//待爬取目标网站如果爬取网络流量正常稳定可以适当减少等待时间
// 500毫秒执行一次请求
var rateLimiter = time.Tick(time.Second / config.Qps)

// fetch到的网页数据 该url不能获取数据则err
func Fetch(url string) ([]byte, error) {

	<-rateLimiter
	log.Printf("Fetching url %s", url)

	//出现浏览器可以对该URL进行访问，可是爬取时则返回403问题
	//resp, err := http.Get(url)
	//defer resp.Body.Close()

	client := &http.Client{}

	//client := getProxyClient()

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("User-Agent",
		"Mozilla/5.0 (Windows NT 6.1; WOW64; Trident/7.0; rv:11.0) like Gecko")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	if err != nil {
		//panic("occured a error")
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil,
			fmt.Errorf("wrong status code %d", resp.StatusCode)
	}

	//判断转义网站编码
	bodyReader := bufio.NewReader(resp.Body)
	encoding := determineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, encoding.NewDecoder())
	return ioutil.ReadAll(utf8Reader)

}

////判断字符编码
//func determineEncoding(r io.Reader) encoding.Encoding{
//	bytes, err := bufio.NewReader(r).Peek(1024)
//	if err != nil {
//		log.Panicf("Fetcher error: %v", err)
//		return unicode.UTF8
//	}
//	encoding, _, _ := charset.DetermineEncoding(bytes, "")
//
//	return encoding
//}

//判断字符编码
func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		// Peek失败 不代表该网页文本不可读 返回默认编码
		log.Panicf("Fetcher error: %v", err)
		return unicode.UTF8
	}
	encoding, _, _ := charset.DetermineEncoding(bytes, "")

	return encoding
}

//func getProxyClient() *http.Client{
//	proxyAddr := "http://125.46.0.62:53281/"
//	//httpUrl := "http://134.175.165.18:8000/get_ip"
//	proxy, err := url.Parse(proxyAddr)
//	if err != nil {
//		log.Fatal(err)
//	}
//	netTransport := &http.Transport{
//		Proxy:http.ProxyURL(proxy),
//		MaxIdleConnsPerHost: 10,
//		ResponseHeaderTimeout: time.Second * time.Duration(5),
//	}
//	httpClient := &http.Client{
//		Timeout: time.Second * 10,
//		Transport: netTransport,
//	}
//
//	return httpClient
//}

//从文件中获取useragent
func getUserAgent() {
	contents, err := ioutil.ReadFile("crawler/fetcher/useragent.txt")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", contents)
}
