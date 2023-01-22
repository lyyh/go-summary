package main

import (
	"context"
	"fmt"
	"time"
)

func request(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, time.Second)

	defer cancel()

	resp := make(chan int)
	go handle(ctx, resp)
	//
	// go (func() {
	// 	time.Sleep(time.Second)
	// 	cancel()
	// })()

	for {
		select {
		// 获取通道值
		case v := <-resp:
			fmt.Println(v)
			// 消息通知 closed。(cancel、timeout、deadline)
		case <-ctx.Done():
			fmt.Println(ctx.Err().Error())
			return
		}
	}

}

func handle(ctx context.Context, resp chan<- int) {
	fmt.Println("handle")
	cache(ctx, resp)
}

func cache(ctx context.Context, resp chan<- int) {
	fmt.Println("cache")
	database(ctx, resp)
}

func database(ctx context.Context, resp chan<- int) {
	// time.Sleep(2 * time.Second)
	go (func() {
		resp <- 10
	})()
	// check done!
	select {
	case <-ctx.Done():
		fmt.Println("ctx2 timeout")
	}
	fmt.Println("database complete")
}

func main() {
	// chain: request -> handle -> cache -> database
	request(context.Background())
	// 接收用户输入，在这里起到了阻塞执行的作用
	fmt.Scanln()
	// output
	// 1. handle
	// 2. cache
	// 10
	// context deadline exceeded
	// ctx2 timeout
	// database complete
}
