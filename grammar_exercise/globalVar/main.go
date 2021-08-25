package main

import (
	"fmt"
	"go_study/grammar_exercise/globalVar/account"
	"go_study/grammar_exercise/globalVar/vars"
)

func main() {
	fmt.Printf("gshadowPath:%s\n", account.GshadowPath)
	vars.Root = "/mnt/"
	fmt.Printf("gshadowPath:%s\n", account.GshadowPath)
	newVal := vars.Root[len(vars.Root)-1:]
	fmt.Println(newVal)
}
