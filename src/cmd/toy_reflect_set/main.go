package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Field1 int
	Field2 string
	Field3 float64
}

func main() {

	s := Sample{1, "F2", 3.0}
	r := reflect.ValueOf(&s).Elem()

	fmt.Println(s)

	for i := 0; i < r.NumField(); i++ {
		f := r.Field(i)
		fmt.Printf("%d: %v %s = %v\n", i, f, f.Type(), f.Interface())

		k := reflect.TypeOf(r.Field(i).Interface()).Kind()

		if k == reflect.Int {
			f.SetInt(100)
		} else if k == reflect.String {
			f.SetString("F2")
		} else if k == reflect.Float64 {
			f.SetFloat(300.0)
		}
	}

	fmt.Println(s)
}
