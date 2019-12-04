package parser

import (
	"Grammar/crawler/engine"
	"regexp"
)

// 正则表达式
// `.+` 中 `.` 表示任意字符 `+` 表示有一个或者多个
// `.*` 中 `.` 表示任意字符 `*` 表示有零个或者多个
// `[0-9a-zA-Z]+` 中 `[0-9a-zA-Z]` 表示字符的规则
// `[^>]` 中 `[^>]` 表示不为 `>` 字符的所有字符
const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-zA-Z]+)" [^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	all := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	limit := 10
	for _, a := range all {
		result.Items = append(result.Items, "City " + string(a[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(a[1]),
			ParserFunc: ParseCity,
		})
		limit --
		if limit == 0 {
			break
		}
	}
	return result
}
