package loop

import (
	"fmt"
	"testing"
)

func TestWhileLoop(t *testing.T) {

	for n := 0; n < 5; {
		fmt.Println(n)
		n++
	}
	a := 0
	// 无限循环
	for {
		if a > 5 {
			break
		}
		fmt.Println(a)
		a++
	}
}
