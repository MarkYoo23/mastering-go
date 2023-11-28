package main

import (
	"fmt"
	"reflect"
)

type Bag struct {
	Tag   string
	Price float64
}

type Human struct {
	Name   string
	Age    int
	OneBag Bag
	Bags   []Bag
}

func main() {
	john := Human{
		Name:   "John",
		Age:    30,
		OneBag: Bag{Tag: "LV", Price: 1000.0},
		Bags: []Bag{
			{Tag: "Robot", Price: 2000.0},
			{Tag: "Pen", Price: 3000.0},
		},
	}

	v := reflect.ValueOf(john)
	findValue(v, 0)
}

func findValue(v reflect.Value, deps int) {

	if v.Kind() == reflect.Struct || v.Kind() == reflect.Slice {
		printStruct(deps, v.Type().Name(), v.Kind())
	} else {
		printField(deps, v.Type().Name(), v, v.Kind())
	}

	for i := 0; i < v.NumField(); i++ {
		findField(v, i, deps+1)
	}
}

func findField(v reflect.Value, i int, deps int) {

	name := v.Type().Field(i).Name
	value := v.Field(i)
	typeKind := v.Field(i).Kind()

	if typeKind == reflect.Struct {
		findValue(value, deps+1)
	} else if typeKind == reflect.Slice {
		for j := 0; j < value.Len(); j++ {
			item := value.Index(j)
			findValue(item, deps+1)
		}
	} else {
		printField(deps, name, value, typeKind)
	}
}

func printStruct(deps int, name string, typeKind reflect.Kind) {
	fmt.Printf("deps %d : %v : %v\n", deps, name, typeKind)
}
func printField(deps int, name string, value reflect.Value, typeKind reflect.Kind) {
	fmt.Printf("deps %d : %v : %v : %v\n", deps, name, value, typeKind)
}
