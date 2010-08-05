package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	name string "namestr"
	age  int
}

func main() {
	p := Person{"Miek Gieben", 33} // No new(), but a composite literal, type of p
	// is now Person, not *Person
	p2 := new(Person)
	p2.name = "Miek Gieben"
	p2.age = 34

	showStruct(p)
	showStruct2(p2)

	fmt.Printf("%v\n", p)
	setStruct(p)
	setStruct(p2)
	fmt.Printf("different? %v\n", p)
}

func showStruct(i interface{}) {
	switch i.(type) {
	case Person:
		fmt.Printf("We are a Person\n")
	default:
		fmt.Printf("We are something else\n")
	}
}

func showStruct2(i interface{}) {
	switch t := i.(type) {
	case *Person:
		r := reflect.NewValue(i)
		tag :=
			// Working with the type here
		r.(*reflect.PtrValue).Elem().Type().(*reflect.StructType).Field(0).Tag
		fmt.Printf("We are a Person, tag %s\n", tag)
	default:
		fmt.Printf("We are something else\n")
	}
}


func setStruct(i interface{}) {
	switch t := i.(type) {
	case Person:
		r := reflect.NewValue(i)
		name :=
		r.(*reflect.StructValue).FieldByName("name").(*reflect.StringValue).Get()
		fmt.Printf("Actual name %s\n", name)
	case *Person:
		r := reflect.NewValue(i)
		name :=
		r.(*reflect.PtrValue).Elem().(*reflect.StructValue).FieldByName("name").(*reflect.StringValue).Get()
		fmt.Printf("Actual name %s\n", name)
		// now set it
		r.(*reflect.PtrValue).Elem().(*reflect.StructValue).FieldByName("name").(*reflect.StringValue).Set("Albert Einstein")
	}
}
