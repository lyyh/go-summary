package for_select

import (
	"fmt"
	"testing"
	"time"
)

func TestForLoop(t *testing.T) {
	i := 0
Loop:
	for {
		select {
		case <-time.After(time.Second * time.Duration(2)):
			i++
			if i == 5 {
				fmt.Println("跳出for循环")
				break Loop
			}
		}
		fmt.Println("for循环内 i=", i)
	}
	fmt.Println("for循环外")
}

func TestGotoLoop(t *testing.T) {
	i := 0
	for {
		select {
		case <-time.After(time.Second * time.Duration(2)):
			i++
			if i == 5 {
				fmt.Println("跳出for循环")
				goto Loop
			}
		}
		fmt.Println("for循环内 i=", i)
	}

Loop:
	fmt.Println("跳到了Loop")
}
