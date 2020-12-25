package sentinelError

import (
	"fmt"
	"github.com/pkg/errors"
	"testing"
)

func TestCountLines(t *testing.T) {
	err := test()
	if err == notFound {
		fmt.Printf("match error\n")
	}

	if err == errors.New("not found") {
		fmt.Printf("don't match error\n") //这条语句不命中
	}
}
