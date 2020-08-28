package type_test

import "testing"
import "fmt"
import "math"

type MyInt int64
// Go语言不允许隐式类型转换
// 别名也不允许隐式类型转换

func	TestImp(t *testing.T){
	var a int32 = 1
	var b int64 = 2
	// b = a
	b = int64(a)
	fmt.Println(a,b)

	var c MyInt = 1234
	var d int64 = 234
	// c = d
	c = MyInt(d)
	fmt.Println(c,d)
	fmt.Println(math.MaxInt64)
	fmt.Println(math.MaxFloat64)
	fmt.Println(math.MaxUint32)
	fmt.Println(math.MaxUint8)
	fmt.Println(math.MaxUint16)
	fmt.Println(math.MaxInt32)
}

func TestPoint(t *testing.T){
	a := 1 
	aPtr := &a
	fmt.Println(*aPtr)
	fmt.Println(aPtr)
	fmt.Printf("%T\n",aPtr)
}

func TestString(t *testing.T){
	var s string
	fmt.Println("*"+s+"*")
	fmt.Println(len(s))
	// fmt.Println(s == nil)
	fmt.Println(s == "")
}