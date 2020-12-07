package main

import (
	"fmt"
	"reflect"

	"github.com/fairyhunter13/reflecthelper/v3"
)

// Get the elem of a variable of reflect.Value
func main() {
	normal()
	usingReflectHelper()
}

func normal() {
	// normally
	var test *int
	x := 5
	test = &x
	elemTest := reflect.ValueOf(&test).Elem().Elem()
	fmt.Println("Elem value is: ", elemTest.Int())
}

func usingReflectHelper() {
	// using reflecthelper
	var test *int
	x := 5
	test = &x
	elemTest := reflecthelper.GetChildElem(reflect.ValueOf(&test))
	fmt.Println("Elem value is: ", elemTest.Int())
}
