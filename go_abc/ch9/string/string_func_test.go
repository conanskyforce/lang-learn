package string_test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestStringFunc(t *testing.T) {
	s := "a,b,c,d,e"
	fmt.Println(s)
	parts := strings.Split(s, ",")
	fmt.Println(parts)
	for _, v := range parts {
		fmt.Println(v)
	}
	d := strings.Join(parts, "-")
	fmt.Println(d)
}

func TestStringConv(t *testing.T) {
	s := "1"
	// 字符串转整型
	d, err := strconv.Atoi(s)
	if err == nil {
		fmt.Println(s, d, err)
	}
	s2 := strconv.Itoa(1234)
	fmt.Println("asd" + s2)
}
