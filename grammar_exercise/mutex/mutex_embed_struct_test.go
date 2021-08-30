package mutex

import (
	"fmt"
	"sync"
	"testing"
)

type TestObj struct {
	sync.Mutex
	data map[string]string
}

//实现实现了Lock和Unlock则会出bug,并没有调用mutex的Lock和Unlock方法来加锁和解锁

func (to *TestObj) Lock() {
	fmt.Println("TestObj Lock function call")
}

func (to *TestObj) Unlock() {
	fmt.Println("TestObj Unlock function call")
}

//TestEmbedMutexInStruct 演示嵌入mutex变量到struct结构体
func TestEmbedMutexInStruct(t *testing.T) {
	obj := TestObj{data: map[string]string{}}
	//加锁
	obj.Lock()

	//写数据
	obj.data["haolipeng"] = "chengdu"

	//解锁
	obj.Unlock()
}
