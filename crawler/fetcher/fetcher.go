package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var rateLimiter = time.Tick(100 * time.Millisecond)

func Fetch(url string) ([]byte, error) {

	<-rateLimiter
	//出现浏览器可以对该URL进行访问，可是爬取时则返回403问题
	//resp, err := http.Get(url)
	//defer resp.Body.Close()

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("User-Agent",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36")

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
		log.Panicf("Fetcher error: %v", err)
		return unicode.UTF8
	}
	encoding, _, _ := charset.DetermineEncoding(bytes, "")

	return encoding
}
