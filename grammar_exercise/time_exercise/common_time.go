package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("time is ", now)

	then := time.Date(2019, 7, 7, 21, 57, 45, 651387237, time.UTC)
	fmt.Println(then)

	fmt.Println("year: ", then.Year())
	fmt.Println("month: ", then.Month())
	fmt.Println("Day: ", then.Day())
	fmt.Println("Hour: ", then.Hour())
	fmt.Println("Minute: ", then.Minute())
	fmt.Println("Second: ", then.Second())
	fmt.Println("Location: ", then.Location())

	diff := now.Sub(then)
	fmt.Println("diff from now is ", diff)
}
