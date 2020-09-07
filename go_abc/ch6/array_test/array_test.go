package array_test

import (
	"fmt"
	"testing"
)

func TestArrayInit(t *testing.T) {
	var arr [3]int
	fmt.Println(arr[0], arr[1])
	var arr1 = [2]int{1, 2}
	fmt.Println(arr1[0], arr1[1])
	// 不指定个数，自动根据初始化个数
	arr2 := [...]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(arr2[0], arr2[6], len(arr2))
}

func TestArrayTravel(t *testing.T) {
	arr3 := [5]int{1, 2, 3}
	// 常规迭代
	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}
	// 迭代器
	for idx, e := range arr3 {
		fmt.Println(idx, e)
	}
	// 多维数组
	arr4 := [3][4]int{{1, 2}, {2, 3}}
	fmt.Println(arr4)
	// 切片
	fmt.Println(arr4[1:2])
	fmt.Println(arr4[0:len(arr4)])
}
