package testsort

import (
	"fmt"
	"testing"
	"unsafe"
)

type Lp struct {
	name string
	age  int
}

func TestPointer(t *testing.T) {
	l := Lp{
		name: "t6e",
		age:  0,
	}
	lPointer := unsafe.Pointer(&l)

	lt := (*string)(unsafe.Pointer(lPointer))
	*lt = "lcl"

	lage := (*int64)(unsafe.Pointer(uintptr(lPointer) + unsafe.Offsetof(l.age)))
	*lage = 22

	fmt.Println("name", l.name, "age", l.age)
}
