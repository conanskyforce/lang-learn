package fib
import (
	"testing"
	"fmt"
)
func TestFibList(t *testing.T){
	var a int = 1
	var b int = 1
	fmt.Print(" ", a)
	for i:=0;i<5;i++{
		fmt.Print(" ", b)
		tmp := a
		a = b
		b = tmp +a
	}
	fmt.Println(" ")
}

func TestExchange(t *testing.T) {
	a := 1
	b := 2
	tmp := a
	a = b
	b = tmp
	fmt.Println(a,b)
	a,b = b,a
	fmt.Println(a,b)
}

const (
	Mon = 1 + iota
	Tues = 2 << iota
	Wen
)

func TestConstant(t *testing.T){
	fmt.Println(Mon,Tues,Wen)
	a := 7 //0111
	fmt.Println(a&Tues == Tues,a&Mon == Mon,a&Wen == Wen)

}