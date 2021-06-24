package mutex

import (
	"sync"
	"testing"
)

//实验证明，拷贝结构体后，mutex的状态被复制过去了，但是mutex不是同一个，
//所以拷贝后，需要把mutext重新初始化，否则会导致非预期的bug

//User 结构体定义
type User struct {
	m        sync.Mutex
	name     string
	age      int
	relation map[string]string
}

func Init() User {
	user := User{
		relation: make(map[string]string),
		name:     "haolipeng",
		age:      31,
	}
	return user
}

func TestMutexInStruct(t *testing.T) {
	user := Init()

	//已加锁
	user.m.Lock()
	t.Logf("origin users mutex pointer:%p mutex:%v\n", &user.m, user.m)

	newUser := user

	//正确写法如下
	newUser.m = sync.Mutex{} //重新初始化结构体中的mutex
	t.Logf("new users mutex pointer:%p mutex:%v\n\n", &newUser.m, newUser.m)

	newUser.m.Lock() //处于已加锁状态的mutex，再次加锁会导致死锁
	newUser.m.Unlock()
	//newUser.m.Unlock() //处于解锁状态的mutex，再次解锁会导致死锁

	//打印出原mutex变量的锁状态
	t.Logf("finally origin mutex pointer:%p mutex:%v\n", &user.m, user.m)

	//实验证明，拷贝结构体后，mutex的状态被复制过去了，但是mutex不是同一个，所以拷贝后，需要把mutext重新初始化，否则会导致非预期的bug
}
