package slice_test

import (
	"fmt"
	"testing"
)

func TestSliceInit(t *testing.T) {
	var s0 []int
	fmt.Println(len(s0), cap(s0))
	s0 = append(s0, 1)
	s0 = append(s0, 1, 2, 3, 4)
	fmt.Println(len(s0), cap(s0))

	// slice 初始化没有指定个数
	s1 := []int{1, 2, 3, 4}
	fmt.Println(len(s1), cap(s1))
	// 初始化的第二种方法, len, cap
	s2 := make([]int, 2, 4)
	fmt.Println(len(s2), cap(s2))
	fmt.Println(
		s2[0], s2[1],
		// s2[2],
	)
}

func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		fmt.Println(len(s), cap(s))
	}
}

func TestSliceShareMemo(t *testing.T) {
	// 共享结构体
	year := []string{
		"jan", "feb", "mar", "apr", "may", "jun", "jul", "aug", "sep", "oct", "nov", "dec",
	}
	Q2 := year[3:6]
	fmt.Println(Q2, len(Q2), cap(Q2))
	summer := year[5:8]
	fmt.Println(summer, len(summer), cap(summer))
	summer[0] = "unknow"
	fmt.Println(Q2)
	fmt.Println(year)
}

func TestSliceComp(t *testing.T) {
	// slice 不能直接比较
	// s1 := []int{}
	// s2 := []int{}
	// if s1 == s2 {
	// 	fmt.Println("s1 == S2")
	// }

}
