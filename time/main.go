package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Weekday())
	fmt.Println(now.Month())
	fmt.Println(now.Year())

	timeInUTC := time.Date(2009, 1, 1, 12, 0, 0, 0, time.UTC)
	fmt.Println(timeInUTC)

	nextTime := now.Add(time.Hour)
	fmt.Println(nextTime)
	sub := nextTime.Sub(now)
	fmt.Println(sub)

	// tick := time.Tick(time.Second)
	// for range tick {
	// 	fmt.Println("1")
	// }

	CST := time.FixedZone("CST", 8*60*60)
	format := now.Format("2006-01-02 15:04:05")
	fmt.Println(format, now.In(CST).Format(time.RFC3339))
}
