package string_test

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	var s string
	fmt.Println(s) // 初始化默认为零值""

	s = "hello"
	// s[1] = 'x' // string 是不可变的slice
	fmt.Println(s, len(s))

	s = "\xE4\xB8\xA5"
	fmt.Println(s, len(s))
	s = "\xE24\xB38\x2A5" // 可以存储任意二进制数据
	fmt.Println(s, len(s))

	s = "中国"
	fmt.Println(s, len(s))

	c := []rune(s)
	fmt.Println(c, len(c))
	fmt.Printf("中国(rune) c[0],c[1] unicode  %x, %x\n", c[0], c[1])
	fmt.Printf("中国 s[0],s[1] UTF8 %x,%x\n", s[0], s[1])
}

func TestStringToRune(t *testing.T) {
	s := "中华人民共和国"
	c := []rune(s)
	fmt.Printf("%c\n", c[0])
	for _, c := range s {
		fmt.Printf("%[1]c,%[1]d,%[1]x\n", c)
	}
}
