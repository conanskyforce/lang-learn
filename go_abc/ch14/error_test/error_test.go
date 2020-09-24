package error_test

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

var LessThanTwoError error = errors.New("n should be bigger than 2")
var BiggerThanHundredError error = errors.New("n should be less than 100")

func GetFib(n int) ([]int, error) {
	if n < 2 {
		return nil, LessThanTwoError
	}
	if n > 100 {
		return nil, BiggerThanHundredError
	}
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}
func TestError(t *testing.T) {
	if v, err := GetFib(10); err == nil {
		fmt.Println(v)
	} else {
		fmt.Println("error:", err)
	}
	if v, err := GetFib(-1); err == nil {
		fmt.Println(v)
	} else {
		fmt.Println("error:", err)
	}
	if v, err := GetFib(101); err == nil {
		fmt.Println(v)
	} else {
		fmt.Println("error:", err)
	}
	GetFib2("23")
}

func GetFib2(str string) {
	var (
		i    int
		err  error
		list []int
	)
	if i, err = strconv.Atoi(str); err != nil {
		fmt.Println("Error", err)
		return
	}
	if list, err = GetFib(i); err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println(list)
}
