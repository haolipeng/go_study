package channel

import "testing"

//测试1：关闭nil channel会触发panic
func TestCloseNilChannel(t *testing.T) {
	var a chan int
	close(a)
}

//测试2：关闭已经处于关闭状态的通道，会引发panic
func TestCloseChannelTwice(t *testing.T) {
	dataCh := make(chan string, 10)
	dataCh <- "hello world"
	close(dataCh)
	close(dataCh) //panic: close of closed channel
}

//测试3：向已关闭通道写入数据
func TestWriteClosedChannel(t *testing.T) {
	dataCh := make(chan string, 10)
	close(dataCh)
	dataCh <- "hello world" //panic: send on closed channel
}
