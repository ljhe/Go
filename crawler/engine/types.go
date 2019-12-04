package engine

// 代码格式化快捷键: CTRL + ALT + L

type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

func NilParse([]byte) ParseResult {
	return ParseResult{}
}
