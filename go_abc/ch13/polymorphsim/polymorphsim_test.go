package polymorphsim

import (
	"fmt"
	"testing"
)

type Code string
type Programmer interface {
	WriteHelloWorld() Code
}

type GoProgrammer struct{}

func (p *GoProgrammer) WriteHelloWorld() Code {
	return "fmt.Println(\"Hello world!\")"
}

type JavaProgrammer struct{}

func (p *JavaProgrammer) WriteHelloWorld() Code {
	return "System.out.Println(\"Hello world!\")"
}

func writeFirstProgram(p Programmer) {
	fmt.Printf("%T, %v \n", p, p.WriteHelloWorld())
}

func TestPoly(t *testing.T) {
	var goProg *GoProgrammer = &GoProgrammer{}
	// GoProgrammer := &GoProgrammer{}
	javaProg := new(JavaProgrammer)
	writeFirstProgram(goProg)
	writeFirstProgram(javaProg)
}
