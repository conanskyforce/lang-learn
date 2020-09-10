package map_ext

import (
	"fmt"
	"testing"
)

func TestMapWithFuncValue(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }
	fmt.Println(m[1](10), m[2](10), m[3](10))
}
func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true
	mySet[2] = true
	delete(mySet, 2)
	IsExists(1, mySet)
	IsExists(2, mySet)

}
func IsExists(key int, set map[int]bool) {
	if set[key] {
		fmt.Println(key, "exists in", set)
	} else {
		fmt.Println(key, "not in", set)
	}
}
