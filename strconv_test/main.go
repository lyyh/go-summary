package main

import (
	"fmt"
	"strconv"
)

func main() {
	atoi, err := strconv.Atoi("123")
	if err == nil {
		fmt.Println(atoi)
	}
	parseBool, err := strconv.ParseBool("T")
	if err == nil {
		fmt.Println(parseBool)
	}
}
