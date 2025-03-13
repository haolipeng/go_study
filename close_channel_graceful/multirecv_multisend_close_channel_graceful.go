package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

//在多个senders，多个receivers的情况下，就需要引入一个中间人，确保只有中间人自己去关闭信号channel stopChan
// 从senders或receivers处发送的关闭dataCh通道的N个请求，中间人只处理一个请求，然后去关闭信号通道 stopCh
//多个receivers监听stopCh上信号，方便退出go协程
//多个senders监听stopCh上信号，方便退出go协程
func main() {
	rand.Seed(time.Now().UnixNano())
	//log.SetFlags(0)

	// ...
	const MaxRandomNumber = 100000
	const NumReceivers = 5
	const NumSenders = 8

	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(NumReceivers)

	// stopCh is an additional signal channel.
	// Its sender is the moderator goroutine shown below. 它的发送者是调解协程
	// Its reveivers are all senders and receivers of dataCh. 它的接收者是dataCh通道的发送者和接收者
	dataCh := make(chan int, 100)
	stopCh := make(chan struct{})

	// 通道用于通知调解者来关闭stopCh通道
	// 通道的发送者是dataCh的发送者和接收者
	// 通道的接收者是调解者协程
	// 通道的容量只有1
	toStop := make(chan string, 1)

	var stoppedBy string

	// moderator 调解人
	go func() {
		stoppedBy = <-toStop
		close(stopCh)
	}()

	// 发送者
	for i := 0; i < NumSenders; i++ {
		go func(id string) {
			defer func() {
				//fmt.Printf("sender(%s) has stopped!\n", id)
			}()
			for {
				value := rand.Intn(MaxRandomNumber)
				if value == 0 {
					// Here, a trick is used to notify the moderator
					// to close the additional signal channel.
					select {
					case toStop <- "sender#" + id:
					default:
					}
					return
				}

				// 第一个select语句是尽可能早的退出go协程. select语句阻塞在接收操作分支和default分支
				// 将会被特别优化为一个尝试接收数据的操作
				select {
				case <-stopCh:
					fmt.Println("触发发送协程1的的stop")
					return
				default:
				}

				// 即使stopCh关闭了, case <-stopCh分支可能仍不会被选中
				// 发送数据到dataCh可能是非阻塞的(比如dataCh为buffered channel)
				// 所以前面的select分支是有必要的
				select {
				case <-stopCh:
					fmt.Println("触发发送协程2的的stop")
					return
				case dataCh <- value:
				}
			}
		}(strconv.Itoa(i))
	}

	// 接收者
	for i := 0; i < NumReceivers; i++ {
		go func(id string) {
			defer func() {
				wgReceivers.Done()
				fmt.Printf("receiver(%s) has stopped!\n", id)
			}()

			for {
				select {
				case <-stopCh:
					fmt.Println("触发接收协程1的的stop")
					return
				default:
				}

				select {
				case <-stopCh:
					fmt.Println("触发接收协程2的的stop")
					return
				case value := <-dataCh:
					//当获取的值是其最大值-1时，尝试关闭接收协程
					if value == MaxRandomNumber-1 {
						// The same trick is used to notify
						// the moderator to close the
						// additional signal channel.
						select {
						case toStop <- "receiver#" + id:
						default:
						}
						return
					}

					//log.Println(value)
				}
			}
		}(strconv.Itoa(i))
	}

	// ...
	wgReceivers.Wait()
	log.Println("stopped by", stoppedBy)
}
