package condition

import (
	"fmt"
	"testing"
)

func TestMultiSec(t *testing.T) {
	if a := 1 == 1; a {
		fmt.Println(a)
	}
	// if v, err := someFunc(); err == nil {
	// 	// no error
	// } else {
	// 	// has error
	// }
}

func TestSwitchCase(t *testing.T) {
	a := 2
	switch {
	case a == 1:
		fmt.Println("a==1")
	case a == 2:
		fmt.Println("a==2")
	}
	switch b := 2; b {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	}
	fmt.Println("*************")
	for i := 0; i < 10; i++ {
		switch i {
		case 1, 2:
			fmt.Println("1,2")
		case 4, 3:
			fmt.Println("3,4")
		default:
			fmt.Println("Not 1,2,3")
		}
		// 像多重 if else 语句
		switch {
		case i%2 == 0:
			fmt.Println(i, "%2 == 0")
		case i%3 == 0:
			fmt.Println(i, "%3 == 0")
		}
	}
}
