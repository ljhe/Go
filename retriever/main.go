package main

import (
	"Grammar/retriever/mock"
	"Grammar/retriever/real"
	"fmt"
)

type Retriever interface {
	Get(url string) string
}

func downLoad(r Retriever)  string {
	return r.Get("http://www.baidu.com")
}

func main()  {
	var r Retriever
	r = mock.Retriever{Contents:"this"}
	r = real.Retriever{
		UserAgent: "",
		TimeOut:   0,
	}
	fmt.Println(downLoad(r))
}
