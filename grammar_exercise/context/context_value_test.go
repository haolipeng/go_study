package context

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestContextValue(t *testing.T) {
	key := "name"
	value := "haolipeng"
	ctx := context.WithValue(context.Background(), key, value)
	go func(ctx context.Context) {
		key := "name"
		value := ctx.Value(key)
		fmt.Println(value)
	}(ctx)

	time.Sleep(time.Second * 4)
}
