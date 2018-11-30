package main

//Men interface
type Men interface {
	sayHi()
	Sing(lyric string)
}

type Human struct {
	name    string
	age     int
	address string
}

type Student struct {
	Human
}

func main() {

}
