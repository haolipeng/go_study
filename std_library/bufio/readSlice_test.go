package bufio

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func TestReadSlice(t *testing.T) {
	s := strings.NewReader("abcdef|ghij")
	r := bufio.NewReader(s)
	token, err := r.ReadSlice('|')
	if err != nil {
		panic(err)
	}
	fmt.Printf("Token: %q\n", token)
}

//没找到分隔符，且缓冲区中
func TestBufferFull(t *testing.T) {
	s := strings.NewReader(strings.Repeat("a", 16) + "|")
	r := bufio.NewReaderSize(s, 16)
	token, err := r.ReadSlice('|')
	if err != nil {
		panic(err)
	}
	fmt.Printf("Token: %q\n", token)
}
