package engine

type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult //url对应Parser解析器
}

type ParseResult struct {
	Requests []Request
	Items    []interface{} //存储任何类型
}
