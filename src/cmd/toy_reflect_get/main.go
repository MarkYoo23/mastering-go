package main

import (
	"fmt"
	"reflect"
)

type Human struct {
	Name string
	Age  int
}

//<main.Human Value>
//Name : John : string
//Age : 30 : int

func main() {
	john := Human{Name: "John", Age: 30}

	v := reflect.ValueOf(john)

	fmt.Println(v.String())

	for i := 0; i < v.NumField(); i++ {

		name := v.Type().Field(i).Name
		value := v.Field(i)
		typeName := v.Field(i).Kind()

		fmt.Printf("%v : %v : %v\n", name, value, typeName)
	}
}
