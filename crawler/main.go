package main

import (
	"Grammar/crawler/engine"
	"Grammar/crawler/zhenai/parser"
)

func main()  {
	engine.Run(engine.Request{
				Url:        "http://www.zhenai.com/zhenghun",
				ParserFunc: parser.ParseCityList,
	})
}