package main

import (
	"errors"
	"fmt"
	"golang.org/x/sync/singleflight"
	"sync"
	"time"
)

type mockClient struct {
	sf singleflight.Group
}

func (client *mockClient) Forget(key string) {
	client.sf.Forget(key)
	return
}

func (client *mockClient) Get(key string) (interface{}, error) {
	fmt.Printf("read %s from database\n", key)
	time.Sleep(2 * time.Second)
	//重新拼装一下key
	v := "Content of key" + key
	return v, nil
}

//同步调用
func (client *mockClient) singleFlightGet(key string) (interface{}, error) {
	fmt.Println("Sync call singleFlightGet function!")
	v, err, _ := client.sf.Do(key, func() (interface{}, error) {
		return client.Get(key)
	})

	if err != nil {
		return nil, err
	}

	return v, err
}

//异步调用
func (client *mockClient) singleFlightGetAsyn(key string) (interface{}, error) {
	fmt.Println("Asyn call singleFlightGetAsyn function!")
	resCh := client.sf.DoChan(key, func() (interface{}, error) {
		return client.Get(key)
	})

	var res singleflight.Result

	timeout := time.After(5 * time.Second)
	select {
	case <-timeout:
		fmt.Println("timeout")
		return nil, errors.New("timeout")
	case res = <-resCh:
		return res.Val, res.Err
	}

	return nil, nil
}

const cnt = 2

func main() {
	client := mockClient{}
	wg := sync.WaitGroup{}

	//10个协程同时获取同一个key的值，到数据库层面实际只调用一次
	for i := 0; i < cnt; i++ {
		wg.Add(1)
		go func(times int) {
			//测试Forget函数
			/*fmt.Printf("times is:%d\n", times)
			if times == 5 {
				//测试
				client.Forget("haolipeng")
			}*/

			// Do版本
			//client.singleFlightGet("haolipeng")

			// DoChan版本
			val, _ := client.singleFlightGetAsyn("haolipeng")
			fmt.Printf("times(%d) val: %s\n", times, val.(string))

			wg.Done()
		}(i)
	}

	wg.Wait()
}
