package main

import "fmt"

func main() {
	s1 := make([]int, 0, 5)
	fmt.Printf("s1切片地址: %p len:%d cap:%d\n", &s1, len(s1), cap(s1))

	//appendFunc(s1)
	appendFuncWithAddress(&s1)

	fmt.Println("s1 slice:", s1)
	fmt.Println("s1 slice expression:", s1[:5])
}

func appendFunc(s2 []int) {
	s2 = append(s2, 1, 2, 3)
	fmt.Printf("s2 slice address %p\n", s2)
	fmt.Println("s2 slice:", s2)
}

func appendFuncWithAddress(s2 *[]int) {
	fmt.Printf("s2切片地址: %p len:%d cap:%d\n", s2, len(*s2), cap(*s2))
	*s2 = append(*s2, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Printf("append后s2切片地址: %p len:%d cap:%d\n", s2, len(*s2), cap(*s2))
	fmt.Println("s2切片: ", *s2)
}
