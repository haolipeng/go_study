package mutex

import (
	"sync"
	"testing"
)

//实验证明，拷贝结构体后，mutex的状态被复制过去了，但是mutex不是同一个，
//拷贝后，把锁的状态也拷贝过去了，所以需要把mutex重新初始化，否则会导致非预期的bug

//A Mutex must not be copied after first use
//https://godoc.org/sync#Mutex

//User 结构体定义
type User struct {
	m        sync.Mutex
	name     string
	age      int
	relation map[string]string
}

//初始化对象，返回结构体
func Init() *User {
	user := &User{
		relation: make(map[string]string),
		name:     "haolipeng",
		age:      31,
	}
	return user
}

func TestMutexInStructBad(t *testing.T) {
	user := Init()

	//已加锁
	user.m.Lock()
	t.Logf("origin users mutex pointer:%p mutex:%v\n", &user.m, user.m)

	//错误写法 error:会将结构体中的mutex锁状态一起复制
	newUser := user

	//打印出拷贝后的mutex变量的锁状态
	t.Logf("copy mutex pointer:%p mutex:%v\n", &user.m, user.m)

	newUser.m.Lock() //处于已加锁状态的mutex，再次加锁会导致死锁
	//执行上句代码后，报错 fatal error: all goroutines are asleep - deadlock!
}

func TestMutexInStructGood(t *testing.T) {
	user := Init()

	//已加锁
	user.m.Lock()
	t.Logf("origin users mutex pointer:%p mutex:%v\n", &user.m, user.m)

	//正确写法如下
	newUser := user
	newUser.m = sync.Mutex{} //重新初始化结构体中的mutex，初始化其锁的状态
	t.Logf("new users mutex pointer:%p mutex:%v\n", &newUser.m, newUser.m)

	t.Log("newUser mutex Unlock")
	newUser.m.Lock() //处于已加锁状态的mutex，再次加锁会导致死锁
	t.Logf("new users mutex pointer:%p mutex:%v\n", &newUser.m, newUser.m)

	t.Log("newUser mutex Unlock")
	newUser.m.Unlock()
	t.Logf("new users mutex pointer:%p mutex:%v\n", &newUser.m, newUser.m)
	//newUser.m.Unlock() //处于解锁状态的mutex，再次解锁会导致死锁

	//实验证明，拷贝结构体后，mutex的状态被复制过去了，但是mutex不是同一个，所以拷贝后，需要把mutext重新初始化，否则会导致非预期的bug
}
