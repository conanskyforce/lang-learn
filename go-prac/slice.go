package main
import (
	"fmt"
)

func main()  {
	slice := make([]int,0,5)
	fmt.Println(slice)
	slice = append(slice, 1, 2, 3,1,124,1,25,15)
	fmt.Println(slice)
	m := map[int]string{
		1: "A",
		2: "B",
		3: "C",
	}
	for k,v := range m {
		fmt.Println(k,v)
		slice = append(slice, k)
	}
	fmt.Println(slice)
	fmt.Println(len(slice),cap(slice))
	slice2 := slice[:2]
	fmt.Println(slice2)
}