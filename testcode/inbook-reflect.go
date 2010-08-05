package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	name string "namestr"
	age int
	}


func main() {
	p1 := new(Person)
	p1.name = "Miek Gieben"
	p1.age  = 33
	ShowTag(p1)
}

func ShowTag(i interface{}) {
	// NEed this because we check for *reflect value
	switch t := reflect.NewValue(i).(type) {
	case *reflect.PtrValue:
		tag := t.Elem().Type().(*reflect.StructType).Field(0).Tag
		fmt.Printf("%s\n", tag)
	}
}
