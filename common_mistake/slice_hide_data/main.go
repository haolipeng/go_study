package main

import "fmt"

func getBadSlice() []byte {
	raw := make([]byte, 10000)
	fmt.Println(len(raw), cap(raw), &raw[0])
	return raw[:3]
}

func getGoodSlice() []byte {
	raw := make([]byte, 10000)
	fmt.Println(len(raw), cap(raw), &raw[0])
	res := make([]byte, 3)
	copy(res, raw[:3])
	return res
}

func main() {
	//test 1
	data := getBadSlice()
	fmt.Println(len(data), cap(data), &data[0])

	fmt.Println("--------------下面是正确的做法---------------------")
	//test 2
	data = getGoodSlice()
	fmt.Println(len(data), cap(data), &data[0])
}
