package extension

import (
	"fmt"
	"testing"
)

type Pet struct{}

func (p *Pet) Speak() {
	fmt.Println("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ", host)
}

type Dog struct {
	// p *Pet
	Pet // 匿名嵌套类型
}

// func (d *Dog) Speak() {
// 	// d.p.Speak()
// 	fmt.Println("Dog Speak")
// }

// func (d *Dog) SpeakTo(host string) {
// 	d.p.SpeakTo(host)
// }

func TestDog(t *testing.T) {
	dog := new(Dog)
	dog.SpeakTo("Chao")
	dog.Speak()

}
