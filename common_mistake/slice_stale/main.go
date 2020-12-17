package main

import "fmt"

//多个slice可以引用同一个数组，比如，当你从一个已有的slice创建一个新的slice时，stale slice就会发生。
//在某些情况下，在一个slice中添加新的数据，在原有数组无法存储更多新的数据时，将导致分配一个新的数组。
//而现在其他的slice还指向老的数组
//参考资料
//https://zhuanlan.zhihu.com/p/148201107
func main() {
	s1 := []int{1, 2, 3}
	fmt.Println(len(s1), cap(s1), s1) //prints 3 3 [1 2 3]

	s2 := s1[1:]
	fmt.Println(len(s2), cap(s2), s2) //prints 2 2 [2 3]

	for i := range s2 {
		s2[i] += 20
	} //prints [22 23]

	//still referencing the same array
	fmt.Println(s1) //prints [1 22 23]
	fmt.Println(s2) //prints [22 23]

	s2 = append(s2, 4)
	for i := range s2 {
		s2[i] += 10
	}

	//s1 is now "stale"
	fmt.Println(s1) //prints [1 22 23]
	fmt.Println(s2) //prints [32 33 14]
}
