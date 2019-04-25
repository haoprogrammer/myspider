package engine

import (
	"log"
)

type SimpleEngine struct {
}

func (e SimpleEngine) Run(seeds ...Request) {
	// Engine 维护Request队列
	var requests []Request

	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		parserResult, err := Worker(r)
		if err != nil {
			continue
		}

		// 添加parseResult所有的Request到requests
		requests = append(requests, parserResult.Requests...)

		for _, item := range parserResult.Items {

			log.Printf("Got Item %v", item)
		}
	}
}
