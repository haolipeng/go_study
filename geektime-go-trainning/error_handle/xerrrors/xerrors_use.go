package main

import (
	"database/sql"
	"fmt"

	"golang.org/x/xerrors"
)

//引入了一个新的 fmt 格式化动词: %w

func bar() error {
	if err := foo(); err != nil {
		return xerrors.Errorf("bar failed: %w", foo())
	}
	return nil
}

func foo() error {
	return xerrors.Errorf("foo failed: %w", sql.ErrNoRows)
}

func main() {
	err := bar()
	if xerrors.Is(err, sql.ErrNoRows) {
		fmt.Printf("data not found, %v\n", err)
		fmt.Printf("%+v\n", err)
		return
	}
	if err != nil {
		// unknown error
	}
}
