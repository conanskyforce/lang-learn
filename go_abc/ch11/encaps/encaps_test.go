package encaps

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	Id   string
	Name string
	Age  int8
}

func TestCreateObj(t *testing.T) {
	conan := Employee{"1", "conan", 28}
	kevin := Employee{Name: "kevin", Age: 36}
	steve := new(Employee)
	fmt.Println(conan, kevin, steve)
	fmt.Println(conan.Name, kevin.Age, steve.Id)
	fmt.Printf("%T,%T,%T\n", conan, &conan, steve)
	fmt.Println(unsafe.Pointer(&conan))
	fmt.Println(conan.String1())

	fmt.Println(unsafe.Pointer(steve))
	fmt.Println(steve.String())
}

func (e Employee) String1() string {
	fmt.Println(unsafe.Pointer(&e))
	return fmt.Sprintf("ID:%s,Name:%s,Age:%d", e.Id, e.Name, e.Age)
}
func (e *Employee) String() string {
	fmt.Println(unsafe.Pointer(e))
	return fmt.Sprintf("ID:%s,Name:%s,Age:%d", e.Id, e.Name, e.Age)
}
