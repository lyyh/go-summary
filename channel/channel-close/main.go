package main

import (
	"fmt"
	"time"
)

func waitClose() {
	closed := make(chan struct{})

	for i := 0; i < 2; i++ {
		go func(i int) {
			select {
			case <-closed:
				fmt.Printf("%d closed\n", i)
			}
		}(i)
	}

	// 发送指令关闭所有协程
	// go的协程不支持直接从外部退出，所以只能通过协程自己退出的方式
	close(closed)

	time.Sleep(time.Second)
}

func waitAfterClose() {
	closed := make(chan struct{})

	for i := 0; i < 2; i++ {
		go func(i int) {
			select {
			case <-closed:
				fmt.Printf("%d closed\n", i)
			}
		}(i)
	}

	// 加个时间条件
	after := time.After(5 * time.Second)
	select {
	case <-after:
		close(closed)
	}

	time.Sleep(time.Second)
}
func main() {
	go waitClose()
	go waitAfterClose()
	time.Sleep(10 * time.Second)
}
