package main

import (
	"context"
)

// const KEY = "trace_id"

// func NewRequestID() string {
// 	// return strings.Replace("abc-123")
// 	return "abc-123"
// }

func NewContextWithTraceID() context.Context {
	ctx := context.WithValue(context.Background(), KEY, NewRequestID())
	return ctx
}

func GetContextValue(ctx context.Context, k string) string {
	s, ok := ctx.Value(k).(string)
	if !ok {
		return ""
	}
	return s
}

// func PrintLog(ctx context.Context, msg string) {
// 	fmt.Printf("%s|info|trace_id=%s|%s", time.Now().Format("2006-01-02 15:04:05"), GetContextValue(ctx, KEY), msg)
// }

func ProcessEnter(ctx context.Context) {
	// PrintLog(ctx, "Golang梦工厂")
}

// func main() {
// 	ProcessEnter(NewContextWithTraceID())
// }
