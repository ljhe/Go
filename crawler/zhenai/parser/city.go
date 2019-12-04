package parser

import (
	"Grammar/crawler/engine"
	"regexp"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)" [^>]*>([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	all := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, a := range all {
		result.Items = append(result.Items, "User " + string(a[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(a[1]),
			ParserFunc: ParseProfile,
		})
	}
	return result
}
