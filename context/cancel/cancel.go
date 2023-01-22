package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	timeout, cancelFunc := context.WithTimeout(context.Background(), time.Second*3)

	defer cancelFunc()

	go Monitor(timeout)

	time.Sleep(time.Second * 5)
}

func Monitor(ctx context.Context) {
	for {
		fmt.Println("Println")
	}
}
