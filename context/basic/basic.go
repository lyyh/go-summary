package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	timeout, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	for {
		select {
		case <-timeout.Done():
			fmt.Println("timeout")
			return
		}
	}
}
