package interface_test

import (
	"fmt"
	"testing"
)

type Programmer interface {
	WriteCode() string
}

type GoProgammer struct {
}

func (gop *GoProgammer) WriteCode() string {
	return "Write Hello World"
}

func TestInterface(t *testing.T) {
	conan := GoProgammer{}
	res := conan.WriteCode()
	fmt.Println(res)
	var kevin Programmer
	kevin = new(GoProgammer)
	fmt.Println(kevin.WriteCode())
}
