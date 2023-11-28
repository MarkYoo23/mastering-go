package main

type Field struct {
	Ptr  uintptr
	Name string
	Type string
}

func (f Field) IsField(ptr uintptr) bool {
	return f.Ptr == ptr
}
