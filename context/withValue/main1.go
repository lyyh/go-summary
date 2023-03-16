package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/gofrs/uuid"
)

const KEY = "traceID"

// 生成 trace_id 随机的 uuid
func NewRequestID() string {
	v1, _ := uuid.NewV1()
	return strings.Replace(v1.String(), "-", "", -1)
}

func NewContextWithRequestID(key string) context.Context {
	return context.WithValue(context.Background(), key, NewRequestID())
}

func GetValueWithContext(ctx context.Context, key string) string {
	s, ok := ctx.Value(key).(string)
	if !ok {
		return ""
	}
	return s
}
func PrintLog(ctx context.Context) {
	fmt.Printf("%s|trace_id|%s:%s", time.Now().Format("2006-01-02 15:04:05"), KEY, GetValueWithContext(ctx, KEY))
}

func main() {
	PrintLog(NewContextWithRequestID(KEY))
	// fmt.Println(NewRequestID())
}
