package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main()  {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		fmt.Println("something error")
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("resp.StatusCode is ", resp.StatusCode)
		return
	}

	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ReadAll error")
		return
	}
	printCityList(all)
}

func printCityList(contents []byte) {
	// 正则表达式
	// `.+` 中 `.` 表示任意字符 `+` 表示有一个或者多个
	// `.*` 中 `.` 表示任意字符 `*` 表示有零个或者多个
	// `[0-9a-zA-Z]+` 中 `[0-9a-zA-Z]` 表示字符的规则
	// `[^>]` 中 `[^>]` 表示不为 `>` 字符的所有字符
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-zA-Z]+)" [^>]*>([^<]+)</a>`)
	all := re.FindAllSubmatch(contents, -1)
	for _, a := range all {
		fmt.Printf("City: %s, URL: %s\n", a[2], a[1])
	}
}