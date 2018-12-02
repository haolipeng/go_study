package main

import (
	"fmt"
	"strconv"
)

type People struct {
	name    string
	age     uint8
	Address Addr
}

type Addr struct {
	city string
}

//验证将数值投递如通道是浅拷贝还是深拷贝
func channelCopyValue() {
	p1 := People{"zhangsan", 26, Addr{"habin"}}
	fmt.Printf("p1(1):%v\n", p1)

	var personChan = make(chan People, 1)

	personChan <- p1

	//外部修改p1的值
	p1.Address.city = "chengdu"
	fmt.Printf("p2(2):%v\n", p1)

	p1_copy := <-personChan
	fmt.Printf("p1_copy:%v\n", p1_copy)
}

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}

	//将值写入channel管道中
	c <- total
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////
var personTotal = 200 //人员总数
var personList []Person = make([]Person, personTotal)
var personCount int //人员数量

type Person struct {
	name    string
	age     uint8
	Address Addr
}
type PersonHandler interface {
	Batch(origs <-chan Person) <-chan Person
	Handle(orig *Person)
}

type PersonHandlerImpl struct {
}

func (handler PersonHandlerImpl) Batch(origs <-chan Person) <-chan Person {
	//0. create dst channel,default capacity is 100
	var dstChan = make(chan Person, 100)

	//1. fetch element from origs
	go func() {
		for p := range origs {
			//2. call user-define Handler function
			handler.Handle(&p)

			//3. pass result to dst channel
			dstChan <- p
		}

		fmt.Println("all the element in origs has been handled!")
		close(dstChan)
	}()

	//4. return dst channel
	return dstChan
}

//user-define handle to modify address
func (handler PersonHandlerImpl) Handle(orig *Person) {
	orig.Address.city += strconv.Itoa(100)
}

func getPersonHandler() PersonHandler {
	return PersonHandlerImpl{}
}

//TODO:函数参数origs表明函数内部只会对通道进行写入操作
func fetchPerson(origs chan<- Person) {
	origsCap := cap(origs)
	buffered := origsCap > 0
	goTicketTotal := origsCap / 2

	go func() {
		p, ok := fetchPerson1()
		if !ok {
			for {
				if !buffered || len(goTicketTotal) == goTicketTotal
			}
		}
	}()
}

//TODO:将处理过的信息进行存储和落盘
func savePerson(dstchan <-chan Person) <-chan byte {

}

func initGoTicket(total int) chan byte {
	var goTicket chan byte
	if 0 == total {
		return goTicket
	}

	//create channel of byte type and init
	goTicket = make(chan byte, total)
	for i := 0; i < total; i++ {
		goTicket <- 1
	}

	return goTicket
}

func fetchPerson1() (Person, bool) {

}

//person数组初始化
func init() {
	for i := 0; i < personTotal; i++ {
		name := fmt.Sprintf("%s%d", "P", i)
		p := Person{name, 32, Addr{"beijing"}}
		personList[i] = p
	}
}

func goroutine_channel() {
	handler := getPersonHandler()
	origs := make(chan Person, 100)
	dstchan := handler.Batch(origs)
	fetchPerson(origs)
	sign := savePerson(dstchan)

	<-sign
}

func test() {
	a := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)

	//将切片传递给协程函数
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)

	x, y := <-c, <-c //receive from c
	fmt.Printf("x(%d) + y(%d) = %d\n", x, y, x+y)
}

func main() {
	//简单测试用例
	test()

	//验证通道中数据的拷贝机制
	channelCopyValue()

	//使用通信方式来共享
	goroutine_channel()
}