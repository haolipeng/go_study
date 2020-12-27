package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.ErrExist
	errWrap := fmt.Errorf("...%w...", err)
}
