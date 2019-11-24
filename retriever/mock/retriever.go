package mock

import "fmt"

type Retriever struct {
	Contents string
}

// 常用接口 相当于 Java 中的 toString
func (r *Retriever) String() string {
	return fmt.Sprintf("Retriver: {Contents=%s}", r.Contents)
}

/**
 * 要改变内容必须使用指针接收者
 * 一致性: 如有指针接收者, 最好都是指针接收者 (这样不会有的是值 有的是指针 容易搞不清楚)
 */

// 实现 Poster 接口中的 Posts 方法
func (r *Retriever) Post(url string, form map[string]string) string {
	r.Contents = form["contents"]
	return r.Contents
}

// 实现 Retriever 接口中的 Get 方法
func (r *Retriever) Get (url string) string {
	return r.Contents
}