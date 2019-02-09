package main

import (
	"regexp"
	"fmt"
)

//@后面是域名gmail.com gmail.net gmail.cn
const text = "My email is ccmouse@gmail.com"

func main() {
	//优化过的正则对象
	regex := regexp.MustCompile(".+@.+\\..+")
	match := regex.FindString(text)
	fmt.Println(match)
}
