package main

import (
	"reflect"
)

type User struct {
	Name string
	Age  int

	fields map[uintptr]Field
}

func NewUser(name string, age int) *User {
	user := &User{Name: name, Age: age, fields: make(map[uintptr]Field)}

	v := reflect.ValueOf(user).Elem()
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		n := t.Field(i).Name
		if n == "fields" {
			continue
		}

		addr := v.Field(i).UnsafeAddr()

		user.fields[addr] = Field{
			Ptr:  addr,
			Name: n,
			Type: t.Field(i).Type.String(),
		}
	}

	return user
}

func (u *User) GetField(ptr uintptr) (Field, bool) {
	f, ok := u.fields[ptr]
	return f, ok
}
