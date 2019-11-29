package main

import (
	"fmt"
	"net/http"
	"runtime"
	"sync"
	"time"
)

func main()  {
	goroutine()
	chanForSleep()
	waitGroupForSleep()
}

/**
 * goroutine 有如下特性:
 * go 的执行是非阻塞的, 不会等待
 * go 后面的函数的返回值会被忽略
 * 调度器不能保证多个 goroutine 的执行次序
 * 没有父子 goroutine 的概念, 所有的 goroutine 是平等地被调度和执行的
 * Go 程序在执行时会单独地为 main 函数创建一个 goroutine, 遇到其他关键字时再去创建其他的 goroutine
 * Go 没有暴漏 goroutine id 给用户, 所以不能在一个 goroutine 里面显示地操作另一个 goroutine, 不过 runtime 包提供了一些
 * 	函数访问和设置 goroutine 的相关信息
 */
func goroutine()  {
	go func() {
		sum := 0
		for i := 0; i < 10000; i++ {
			sum += i
		}
		fmt.Println(sum)
		time.Sleep(1 * time.Second)
	}()

	// NumGoroutine 可以返回当前程序的 goroutine 数目
	fmt.Println("NumGoroutine=", runtime.NumGoroutine())

	// main goroutine 故意 "sleep" 5秒, 防止提前退出
	// goroutine 是并发执行的
	// 这里如果不 sleep 那么等不到输出 sum 的结果, 程序就执行完了
	time.Sleep(5 * time.Second)
}

/**
 * goroutine 运行结束后退出, 写到缓存通道中的数据不会消失, 它可以缓冲和适配两个 goroutine 处理速率不一致的情况
 */
func chanForSleep()  {
	c := make(chan struct{})
	go func(c chan struct{}) {
		sum := 0
		for i := 0; i < 10000; i++ {
			sum += i
		}
		fmt.Println(sum)
		// 写通道
		c <- struct{}{}
	}(c)

	// NumGoroutine 可以返回当前程序的 goroutine 数目
	fmt.Println("NumGoroutine=", runtime.NumGoroutine())
	// 读通道 c, 通过通道进行同步等待
	// 使用无缓冲的通道来实现 goroutine 之间的同步等待
	// 此时 c 没有任何元素值 会被阻塞在这里
	<-c
}

/**
 * WaitGroup 用来等待多个 goroutine 完成
 * 调用 Add 设置需要等待的数目
 * 调用 Done 来表示完成一个 goroutine
 * Wait 被用来等待所有的 goroutine 完成
 */
func waitGroupForSleep()  {
	var wg sync.WaitGroup
	var urls = []string{
		"http://www.qq.com",
		"http://www.baidu.com",
		"https://www.jiulingwan.com",
	}

	for _, url := range urls {
		fmt.Println("first")
		// 每一个 URL 启动一个 goroutine, 同时给 wg 加 1
		wg.Add(1)

		go func(url string) {
			fmt.Println("third")
			// 当前 goroutine 结束后给 wg 计数减 1, wg.Done() 等价于 wg.Add(-1)
			//defer wg.Add(-1)
			defer wg.Done()

			resp, err := http.Get(url)
			if err == nil {
				fmt.Println(resp.Status)
			}
		}(url)
		fmt.Println("second")
	}

	wg.Wait()
	fmt.Println("the end")

	// 程序运行结果
	//first
	//second
	//first
	//second
	//first
	//second
	//third
	//third
	//third
	//200 OK
	//200 OK
	//200 OK
	//the end
}
