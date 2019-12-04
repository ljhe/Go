package main

import (
	"fmt"
	"math/rand"
)

func GeneratorIntA(done chan struct{}) chan int {
	ch := make(chan int, 5)
	go func() {
	Label:
		for  {
			select {
			case ch <- rand.Int():
			case <-done:
				break Label
			}
		}
		close(ch)
	}()
	return ch
}

func GeneratorIntB(done chan struct{}) chan int {
	ch := make(chan int, 10)
	go func() {
	Label:
		for  {
			select {
			case ch <- rand.Int():
			case <-done:
				break Label
			}
		}
		close(ch)
	}()
	return ch
}

func GeneratorInt(done chan struct{}) chan int {
	ch := make(chan int)
	send := make(chan struct{})
	go func() {
	Label:
		for  {
			select {
			case ch <- <-GeneratorIntA(send):
			case ch <- <-GeneratorIntB(send):
			case <-done:
				send <- struct{}{}
				send <- struct{}{}
				break Label
			}
		}
		close(ch)
	}()
	return ch
}

func main()  {
	// 创建一个作为接受者退出信号的 chan
	done := make(chan struct{})

	// 启动生成器
	ch := GeneratorInt(done)

	// 获取生成器资源
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}

	// 通知生产者停止生产
	// done 通道在这里变成可读的 select 中就有机会走到 case <-done 这个分支中
	done <- struct{}{}
	fmt.Println("stop generator")
}
