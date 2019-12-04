package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main()  {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServer:", err)
	}
}

func sayHello(w http.ResponseWriter, r *http.Request)  {
	// 解析参数 默认不会解析
	err := r.ParseForm()
	if err != nil {
		log.Fatal("ParseForm:", err)
	}
	fmt.Println(r.Form) // 这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	_, err = fmt.Fprintf(w, "Hello astaxie!") // 这个写入到w的是输出到客户端的
	if err != nil {
		log.Fatal("fmt.Fprintf:", err)
	}
}
