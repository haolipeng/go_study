package main

import (
	"fmt"
	"sync"
)

//https://mozillazg.com/2019/04/notes-about-go-lock-mutex.html
//拷贝结构体时可能导致非预期的死锁
//1.mutex作为函数参数时，是传值还是指针？
//编译器会进行报错 Warning:(27, 20) Call of 'useMutexForUnlock' copies lock value: type 'sync.Mutex' is 'sync.Locker'
func useMutexForUnlock(mu sync.Mutex) {
	//打印出来地址以及内容值
	fmt.Printf("before call useMutexForUnlock function,%p %v \n", &mu, mu) //{1 0}
	mu.Unlock()
	fmt.Printf("after call useMutexForUnlock function,%p %v \n\n", &mu, mu)
}

func testMutexAsFuncArgs() {
	var mu sync.Mutex
	fmt.Printf("init mutex: %p %v\n", &mu, mu)
	mu.Lock()
	fmt.Printf("after Lock mutex: %p %v\n\n", &mu, mu)

	//在传如useMutex函数时,mu状态处于加锁状态
	useMutexForUnlock(mu)
	useMutexForUnlock(mu)

	//在函数内部调用useMutexForUnlock解锁，操作的仅是一份拷贝，所以mu的锁还是1
	fmt.Printf("after useMutexForUnlock() mutex: %p %v\n\n", &mu, mu)
}

type User struct {
	m        sync.Mutex
	name     string
	age      int
	relation map[string]string
}

func testMutexInStruct() {
	user := User{
		relation: make(map[string]string),
		name:     "haolipeng",
		age:      31,
	}
	//已加锁
	user.m.Lock()

	fmt.Printf("origin users:%v mutex:%p\n", user, &user.m)

	newUser := user
	//正确写法
	newUser.m = sync.Mutex{} //重新初始化结构体中的mutex
	fmt.Printf("new users:%v mutex:%p\n\n", newUser, &newUser.m)
	newUser.m.Lock() //处于已加锁状态的mutex，再次加锁会导致死锁
	newUser.m.Unlock()
	//newUser.m.Unlock() //处于解锁状态的mutex，再次解锁会导致死锁

	//origin users:{{1 0} haolipeng 31 map[]} mutex:0xc0000b6330
	//new users:{{1 0} haolipeng 31 map[]} mutex:0xc0000b63f0
	//实验证明，拷贝结构体后，mutex的状态被复制过去了，但是mutex不是同一个，所以拷贝后，需要把mutext重新初始化，否则会导致非预期的bug
}

func main() {
	testMutexAsFuncArgs()
	testMutexInStruct()
}
