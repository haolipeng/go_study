package context

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func getContextValue(ctx context.Context) {
	key := "name"
	value := ctx.Value(key)
	fmt.Printf("value in context:%v\n", value)
}

func TestContextValue(t *testing.T) {
	//传递任意类型的key和value都可以
	key := "name"
	value := 123
	ctx := context.WithValue(context.Background(), key, value)
	go getContextValue(ctx)

	time.Sleep(time.Second * 4)
}
