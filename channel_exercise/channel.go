package main

import (
	"fmt"
	"strconv"
	"os"
)

type People struct {
	name    string
	age     uint8
	Address Addr
}

type Addr struct {
	city string
}

/////////////////////////////////验证将数值投递如通道是浅拷贝还是深拷贝///////////////////////////////
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
var personCount int //全局的人员数量

type Person struct {
	name    string
	age     int
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
	//add number 100 to address.city tail
	orig.Address.city += strconv.Itoa(100)
}

//获取PersonHandler接口的实现类型
func getPersonHandler() PersonHandler {
	return PersonHandlerImpl{}
}

//函数参数origs表明函数内部只会对通道进行写入操作
func fetchPerson(origs chan<- Person) {
	origsCap := cap(origs)
	buffered := origsCap > 0
	goTicketTotal := origsCap / 2
	goTicket := initGoTicket(goTicketTotal)

	go func() {
		for {
			p, ok := fetchPerson1()
			if !ok {
				//close channel,notify the receiver of channel
				fmt.Println("all the infomation has been fetched!")
				close(origs)
				break
			}

			if buffered {
				<-goTicket
				go func() {
					origs <- p
					goTicket <- 1
				}()
			} else {
				origs <- p
			}
		}
	}()
}

//将处理过的信息进行存储和落盘
func savePerson(dstchan <-chan Person) <-chan byte {
	sign := make(chan byte, 1)

	//创建文件来保存信息
	outputHandle, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("create file output.txt failed")
	}
	go func() {
		//可以用for range 遍历channel通道
		/*
		for p:=range dstchan{
			//call internal saveperson function
			savePersonInfoInternal(p,outputHandle)
		}
		fmt.Println("All the information has been saved.")
		sign<-0
		*/

		for {
			p, ok := <-dstchan
			if !ok {
				fmt.Println("All the information has been saved.")
				sign <- 0
				break
			}

			//call internal saveperson function
			savePersonInfoInternal(p, outputHandle)
		}
	}()

	return sign
}

func savePersonInfoInternal(p Person, hFile *os.File) {
	strTotal := p.name + strconv.Itoa(p.age) + p.Address.city + "\r\n"
	hFile.WriteString(strTotal)
}

//channal byte 作用是什么
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
	if personCount < personTotal {
		p := personList[personCount]
		personCount++

		return p, true
	}

	return Person{}, false
}

//main函数之前person数组初始化
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
	//简单测试
	test()

	//验证通道中数据的拷贝机制
	channelCopyValue()

	//使用通信方式来共享
	goroutine_channel()
}