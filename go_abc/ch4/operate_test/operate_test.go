package operate_test

import (
	"testing"
	"fmt"
)
func TestComp(t *testing.T){
	a := [...]int{1,2,3}
	b := [...]int{4,5,6}
	fmt.Println(a,b)
	fmt.Println(a == b)
	a[0] = 4
	a[1] = 5
	a[2] = 6
	fmt.Println(a == b)
}

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestBitClear(t *testing.T){
	a := 7 // 0111
	fmt.Println(
		a&Readable == Readable,
		a&Writable == Writable,
		a&Executable == Executable,
	)
	// 按位清零
	// 右边为1，左边清零
	// 右边为0，左边保留
	// a = a &^ Readable
	// a = a &^ Writable
	a = a &^ Executable
	fmt.Println(
		a&Readable == Readable,
		a&Writable == Writable,
		a&Executable == Executable,
	)
}