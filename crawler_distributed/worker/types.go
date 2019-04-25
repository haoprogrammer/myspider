package worker

import (
	"fmt"
	"github.com/pkg/errors"
	"haoprogrammer/myspider/crawler/engine"
	"haoprogrammer/myspider/crawler/zhenai/parser"
	"haoprogrammer/myspider/crawler_distributed/config"
	"log"
)

//函数的序列化
type SerializedParser struct {
	Name string //函数名
	Args interface{}
}

type Request struct {
	Url    string
	Parser SerializedParser
}

type ParseResult struct {
	Items    []engine.Item
	Requests []Request
}

func SerializeRequest(r engine.Request) Request {
	name, args := r.Parser.Serialize()
	return Request{
		Url: r.Url,
		Parser: SerializedParser{
			Name: name,
			Args: args,
		},
	}
}

func SerializeResult(r engine.ParserResult) ParseResult {
	result := ParseResult{
		Items: r.Items,
	}

	for _, req := range r.Requests {
		result.Requests = append(result.Requests, SerializeRequest(req))
	}

	return result
}

func DeserializeRequest(r Request) (engine.Request, error) {
	parser, err := deserializeParser(r.Parser)
	if err != nil {
		return engine.Request{}, err
	}

	return engine.Request{
		Url:    r.Url,
		Parser: parser,
	}, nil
}

func DeserializeResult(r ParseResult) engine.ParserResult {
	result := engine.ParserResult{
		Items: r.Items,
	}

	for _, req := range r.Requests {
		engineReq, err := DeserializeRequest(req)
		if err != nil {
			log.Printf("error deserializing request:%v", err)
			continue
		}

		result.Requests = append(result.Requests, engineReq)
	}
	return result
}

//从serializeParser转换成能被系统识别的parser
//把字符串转换为对应函数
func deserializeParser(p SerializedParser) (engine.Parser, error) {
	switch p.Name {
	case config.ParseCityList:
		return engine.NewFuncParser(
			parser.ParseCityList, config.ParseCityList), nil
	case config.ParseCity:
		return engine.NewFuncParser(
			parser.ParseCity, config.ParseCity), nil
	case config.NilParser:
		return engine.NilParser{}, nil
	case config.ParseProfile:
		if userName, ok := p.Args.(string); ok {
			return parser.NewProfileParser(
				userName), nil
		} else {
			return nil, fmt.Errorf("invalid args:%v", p.Args)
		}
	default:
		return nil, errors.New("unknown parser name")
	}
}

// 序列化后类似
// {"ParserCityList","nil"} , {"ProfileParser",userName}

//将worker包装成RPC函数
