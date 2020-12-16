package main

import "log"

/*
func Fatal(v ...interface{}) {
	std.Output(2, fmt.Sprint(v...))
	os.Exit(1)
}
*/
//log.Fatal函数中已经包含os.Exit(1)
func main() {
	log.Fatalln("Fatal Level: log entry") //app exits here
	log.Println("Normal Level: log entry")
}
