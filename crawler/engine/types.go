package engine

type ParserFunc func(contents []byte, url string) ParserResult

type Request struct {
	Url        string
	ParserFunc ParserFunc
}

type ParserResult struct {
	Requests []Request
	// interface{}是任何类型
	//Items []interface{}
	Items []Item
}

type Item struct {
	Url     string
	Type    string
	Id      string
	Payload interface{}
}

func NilParser([]byte) ParserResult {
	return ParserResult{}
}
