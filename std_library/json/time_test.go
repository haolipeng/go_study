package jsonTest

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

type Info struct {
	StartTime time.Time `json:"StartTime"`
}

func TestJsonTimeFormat(t *testing.T) {
	startTime := time.Now()
	info := Info{StartTime: startTime}
	bytes, err := json.Marshal(info)
	if err != nil {
		panic(0)
	}
	fmt.Println("value:", string(bytes))
}
