package main

import (
	"fmt"
	"reflect"
)

func main()  {
	str1 := "Go语言真香"
	str2 := "解构真香"
	fmt.Println(reflect.TypeOf(str1[3]).Kind())
	fmt.Println(str1[2],string(str1[2]))
	fmt.Printf("%d %c \n", str1[0],str1[1])
	fmt.Println("len(str2):", len(str2))
	
	str3 := []rune(str2)
	fmt.Println(str3[0],str3[1],str3[2],str3[3])
	fmt.Println(string(str3[0]),string(str3[1]),string(str3[2]),string(str3[3]))
}