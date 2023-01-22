package main

import (
	"fmt"
	"time"
)

var ch = make(chan int)

func main() {
	fmt.Println(time.Now().Unix())
	select {
	case <-ch:
		fmt.Println("aaaa")
	case <-time.After(time.Second * 10):
		fmt.Println(3000)
	}
}
