package worker

import "haoprogrammer/myspider/crawler/engine"

type CrawService struct{}

func (CrawService) process(
	req Request, result *ParseResult) error {
	engineReq, err := DeserializeRequest(req)
	if err != nil {
		return err
	}

	engineResult, err := engine.Worker(engineReq)
	if err != nil {
		return err
	}

	*result = SerializeResult(engineResult)
	return nil

}
