package main

import (
	"Grammar/retriever/mock"
	"Grammar/retriever/real"
	"fmt"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

// 接口的组合, RetrieverPoster 里既有 Retriever 又有 Poster
type RetrieverPoster interface {
	Retriever
	Poster
}

func downLoad(r Retriever)  string {
	return r.Get("http://www.baidu.com")
}

func post(poster Poster)  {
	poster.Post("http://www.baidu.com", map[string]string{
		"a" : "x",
		"b" : "y",
	})
}

// 既使用了 Poster 接口中的 Post 方法, 又使用了 Retriever 接口中的 Get 方法
func session(s RetrieverPoster)  string {
	s.Post("http://www.baidu.com", map[string]string{
		"contents": "that",
	})
	return s.Get("http://www.baidu.com")
}

func main()  {
	var r Retriever
	r = &mock.Retriever{Contents:"this"}
	fmt.Println(r)
	r = real.Retriever{
		UserAgent: "",
		TimeOut:   0,
	}
	fmt.Println(downLoad(r))
	s := &mock.Retriever{Contents:"this"}
	fmt.Println(session(s))
}
