package main

import (
	"fmt"
	"os"
	"os/exec"
)

//TODO:需持续补充下
func main() {
	fmt.Println("use cmd example")
	path, _ := os.Getwd()
	path = path + "\\cmd.exe"
	cmd := exec.Command(path)
	res, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("error accoured:", err.Error())
	}
	fmt.Printf("result is %s\n", string(res))
}
