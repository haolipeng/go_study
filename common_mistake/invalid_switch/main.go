package main

import "fmt"

func main() {
	isSpace := func(ch byte) bool {
		switch ch {
		case ' ':
		case '\t':
			return true
		}
		return false
	}
	fmt.Println(isSpace('\t')) //true
	fmt.Println(isSpace(' '))  //false
}
