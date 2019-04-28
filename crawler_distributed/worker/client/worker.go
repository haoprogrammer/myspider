package client

import (
	"haoprogrammer/myspider/crawler/engine"
	"haoprogrammer/myspider/crawler_distributed/config"
	"haoprogrammer/myspider/crawler_distributed/worker"
	"net/rpc"
)

//做转换，通过rpc调用远程方法进行processor
//传统面向对象编程里面会建立 clients  []*rpc.Client  对比下 使用  go语言 chan的好处
func CreateProcessor(clientChan chan *rpc.Client) engine.Processor {
	//client, err := rpcsupport.NewClient(fmt.Sprintf(":%d", config.WorkerPort0))
	//if err != nil {
	//	return nil, err
	//}

	return func(req engine.Request) (engine.ParserResult, error) {
		sReq := worker.SerializeRequest(req)
		var sResult worker.ParseResult
		c := <-clientChan
		err := c.Call(config.CrawlServiceRpc, sReq, &sResult)
		if err != nil {
			return engine.ParserResult{}, nil
		}
		return worker.DeserializeResult(sResult), nil
	}

}
