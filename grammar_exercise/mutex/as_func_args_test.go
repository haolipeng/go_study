package mutex

import (
	"fmt"
	"sync"
	"testing"
)

//https://mozillazg.com/2019/04/notes-about-go-lock-mutex.html
//拷贝结构体时可能导致非预期的死锁
//1.mutex作为函数参数时，是传值
//编译器会进行报错 Warning:(27, 20) Call of 'useMutexForUnlock' copies lock value: type 'sync.Mutex' is 'sync.Locker'
func useMutexForUnlock(mu sync.Mutex) {
	//打印出来地址以及内容值
	fmt.Printf("before call useMutexForUnlock function,%p %v \n", &mu, mu) //{1 0}
	mu.Unlock()
	fmt.Printf("after call useMutexForUnlock function,%p %v \n\n", &mu, mu)
}

func TestFuncArgs(t *testing.T) {
	var mu sync.Mutex
	fmt.Printf("init mutex: %p %v\n", &mu, mu)
	mu.Lock()
	fmt.Printf("after Lock mutex: %p %v\n\n", &mu, mu)

	//mu状态处于加锁状态,虽然在useMutexForUnlock函数进行了解锁，但是操作的是复制的那份，原来的mutex变量的状态并未改变
	useMutexForUnlock(mu)

	//在函数内部调用useMutexForUnlock解锁，操作的仅是一份拷贝，所以mu的锁还是1
	fmt.Printf("after useMutexForUnlock() mutex: %p %v\n\n", &mu, mu)
}
