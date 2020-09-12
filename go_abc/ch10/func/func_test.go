package func_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMulti() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}
func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		end := time.Now()
		fmt.Println(end, start)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}
func timeTest(op int) int {
	a := op
	for a < 1e3 {
		a++
	}
	time.Sleep(time.Second * 1)
	return a
}
func TestFuncMultiReturn(t *testing.T) {
	a, b := returnMulti()
	fmt.Println(a, b)
	fmt.Println(returnMulti())
	ts := timeSpent(timeTest)
	fmt.Println(ts(10))
}

func Sum(op ...int) int {
	var sum int
	for _, v := range op {
		sum += v
	}
	return sum
}
func TestVarParams(t *testing.T) {
	fmt.Println(Sum(1, 2, 3, 4, 5))
}

func TestDeferFunc(t *testing.T) {
	defer fmt.Println(Sum(10))
	defer func() {
		time.Sleep(time.Second * 2)
		fmt.Println("defer to clearn resources")
	}()
	fmt.Println(123)
	// panic(1234)
}
