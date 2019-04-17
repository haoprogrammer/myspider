package engine

type Request struct {
	Url        string
	ParserFunc func([]byte) ParserResult
}

type ParserResult struct {
	Requests []Request
	// interface{}是任何类型
	Items []interface{}
}

func NilParser([]byte) ParserResult {
	return ParserResult{}
}
