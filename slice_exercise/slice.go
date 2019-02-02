package main

import "fmt"

func updateSlices(s []int) {
	s[0] = 100
}

func main() {
	//1.初始化数组
	arr := [...] int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("arr[2:6]", arr[2:6])
	fmt.Println("arr[2:]", arr[2:])
	fmt.Println("arr[:6]", arr[:6])
	fmt.Println("arr[:]", arr[:])

	//2.slice本身无数据，是对其底层数组array的一个view
	s := arr[2:6]

	//3.传参改变slice元素，slice集合是传引用类型,使用后还原0索引上的元素为0
	fmt.Println("after updateSlices function")
	updateSlices(s)
	fmt.Println(s)
	s[0] = 2

	//4.slice的扩展,向后扩展
	//outout:
	//first is  [2 3 4 5]
	//second is  [5 6 7]
	fmt.Println("Slices Extend")
	first := arr[2:6]
	second := first[3:6]
	fmt.Println("first is ", first)
	fmt.Println("second is ", second)

	//slice可以向后扩展，但是不可以向前扩展
	//s[i]不可以超越len(s),向后扩展不可以超越底层数组cap(s)
	//好处，通过slice扩展造成数据泄漏，也不会发生数组越界的情况

	//5.slice的append,当append添加的元素超过了其底层数组的len(s)时
	//系统会申请一个新的array，并把旧元素都拷贝过去
	s1 := arr[2:6]
	s2 := s1[3:5]
	fmt.Printf("s2 underlying arrar address is %p\n", s2)
	s3 := append(s2, 10)
	fmt.Printf("s3 underlying arrar address is %p\n", s3)
	s4 := append(s3, 11)
	fmt.Printf("s4 underlying arrar address is %p\n", s4)
	s5 := append(s4, 12)
	fmt.Printf("s5 underlying arrar address is %p\n", s5)

	fmt.Println("s3 is ", s3)
	fmt.Println("s4 is ", s4)
	fmt.Println("s5 is ", s5)
	fmt.Println("arr array is ", arr)
	//s4,s5 no long view arr array

	//6.cap的变化为2的指数增长
	var appendArray []int
	for i := 0; i < 100; i++ {
		printLenCap(appendArray)
		appendArray = append(appendArray, i)
	}
	fmt.Println(appendArray)

	//7.用make来创建slice,可显十指定len和cap
	s6 := make([]int, 16)
	s7 := make([]int, 10, 32)
	printLenCap(s6)
	printLenCap(s7)

	//8.删除元素(append来删除slice中多余元素)
	fmt.Println("befor Deleting elements from slices")
	fmt.Println(s1)
	s1 = append(s1[:3], s[4:]...)
	fmt.Println("after Deleting elements from slices")
	fmt.Println(s1)
}

func printLenCap(s []int) {
	fmt.Printf("len=%d cap=%d\n", len(s), cap(s))
}
