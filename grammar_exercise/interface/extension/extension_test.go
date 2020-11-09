package extension

//golang的扩展和复用

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Print("Pet Speak() called")
}

func (p *Pet) SpeakTo(name string) {
	p.Speak()
	fmt.Println(" ", name)
}

type Dog struct {
	Pet
}

/*func (d *Dog) Speak() {
	fmt.Println("dog Speak() called")
}*/

func (d *Dog) SpeakTo(name string) {
	d.Speak()
	fmt.Println(" ", name)
}

func TestExtension(t *testing.T) {
	//dog := new(Dog)
	//var dog Pet = new(Dog) //cannot use new(Dog) (type *Dog) as type Pet in assignment

	var dog *Dog = new(Dog)
	dog.Speak()
	dog.SpeakTo("Chao")
}
