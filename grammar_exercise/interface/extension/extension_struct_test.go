package extension

//golang的扩展和复用

import (
	"fmt"
	"testing"
)

type Pet struct {
	name string
}

func (p *Pet) Speak() {
	fmt.Print("Pet Speak() called\n")
}

func (p *Pet) SpeakTo(name string) {
	p.Speak()
	fmt.Println(" ", name)
}

type Dog struct {
	Pet
	age int
}

//注释掉Dog的Speak()函数，则调用Pet的Speak()函数
//func (d *Dog) Speak() {
//	fmt.Println("dog Speak() called")
//}

func (d *Dog) SpeakTo(name string) {
	d.Speak()
	fmt.Println(" ", name)
}

func TestExtension(t *testing.T) {
	//dog := new(Dog)
	//var dog Pet = new(Dog) //cannot use new(Dog) (type *Dog) as type Pet in assignment
	//结构体嵌入方式的初始化方式
	chuba := Dog{
		Pet: Pet{name: "初八"},
		age: 4,
	}
	chuba.Speak()

	var dog *Dog = new(Dog)
	dog.Speak()
	dog.SpeakTo("Chao")
}
