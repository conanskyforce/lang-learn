package empty_interface

import (
	"fmt"
	"testing"
)

func DoSomething(p interface{}) {
	if i, ok := p.(int); ok {
		fmt.Println("Integer", i)
		return
	}
	if i, ok := p.(string); ok {
		fmt.Println("String", i)
		return
	}
	fmt.Printf("Unknow type: %T\n", p)
}

func TestEmptyInterface(t *testing.T) {
	DoSomething(12)
	DoSomething("asf")
	DoSomething('\'')
	DoSomething(0x24)
	DoSomething(123.23)
	DoSomething(make([]int, 2))
	DoSomething([]string{})
}
