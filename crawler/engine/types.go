package engine

type ParserFunc func(contents []byte, url string) ParserResult

//定义Parser接口
type Parser interface {
	Parse(contents []byte, url string) ParserResult
	Serialize() (name string, args interface{})
}

type Request struct {
	Url string
	//ParserFunc ParserFunc
	Parser Parser
}

//函数的序列化
//type SerializedParser struct {
//	Name string		//函数名
//	Args interface{}
//
//}

// 序列化后类似
// {"ParserCityList","nil"} , {"ProfileParser",userName}

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

//func NilParser([]byte) ParserResult {
//	return ParserResult{}
//}
type NilParser struct{}

func (NilParser) Parse(contents []byte, url string) ParserResult {
	return ParserResult{}
}

func (NilParser) Serialize() (name string, args interface{}) {
	return "NilParser", nil
}

type FuncParser struct {
	parser ParserFunc
	name   string
}

func (f *FuncParser) Parse(contents []byte, url string) ParserResult {
	//包含的是函数ParserFunc
	return f.parser(contents, url)
}

func (f *FuncParser) Serialize() (name string, args interface{}) {
	return f.name, nil
}

//工厂方法创建
func NewFuncParser(p ParserFunc, name string) *FuncParser {
	return &FuncParser{
		parser: p,
		name:   name,
	}

}
