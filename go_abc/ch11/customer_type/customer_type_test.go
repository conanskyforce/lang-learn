package customer_type

import (
	"fmt"
	"testing"
	"time"
)

type IntConv func(op int) int

// func timeSpent(inner func(op int) int) func(op int) int {
// 	return func(n int) int {
// 		start := time.Now()
// 		ret := inner(n)
// 		end := time.Now()
// 		fmt.Println(end, start)
// 		fmt.Println("time spent:", time.Since(start).Seconds())
// 		return ret
// 	}
// }
func timeSpent(inner IntConv) IntConv {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		end := time.Now()
		fmt.Println(end, start)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}
func slow(op int) int {
	time.Sleep(time.Second)
	return op
}

func TestCustomerType(t *testing.T) {
	timeSlow := timeSpent(slow)
	fmt.Println(timeSlow(1))
}
