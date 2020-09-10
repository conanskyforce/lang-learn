package map_test

import (
	"fmt"
	"testing"
)

func TestMapInit(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}

	m2 := map[int]int{}
	m2[1] = 1234

	m3 := make(map[string]int, 1)
	m3["asd"] = 1234
	m3["423"] = 1234
	m3["124"] = 1234
	fmt.Println(m, len(m), m["a"])
	fmt.Println(m2, len(m2), m2[1])
	fmt.Println(m3, len(m3))
}

func TestAccessNotExistKet(t *testing.T) {
	m1 := map[string]int{}
	fmt.Println(m1["1"])
	m1["2"] = 0
	fmt.Println(m1["2"])
	if v, ok := m1["2"]; ok {
		fmt.Println("m1[\"2\"] exists and == ", v)
	}
}
func TestMapTravel(t *testing.T) {
	m1 := map[string]string{
		"name":   "Jhon conan",
		"age":    "12",
		"gender": "male",
	}
	for k, v := range m1 {
		fmt.Println(k, v)
	}
}
