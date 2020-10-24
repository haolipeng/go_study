package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()

	millis := nanos / 1000000
	fmt.Println(secs)   //10位
	fmt.Println(millis) //13位
	fmt.Println(nanos)  //19位

	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))
}
