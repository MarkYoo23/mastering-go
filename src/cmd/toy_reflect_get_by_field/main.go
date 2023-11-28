package main

import (
	"fmt"
	"unsafe"
)

func main() {
	user := NewUser("John", 30)

	ageField, ok := user.GetField(uintptr(unsafe.Pointer(&user.Age)))
	if !ok {
		panic("field not found")
	}

	fmt.Println(ageField)
}
