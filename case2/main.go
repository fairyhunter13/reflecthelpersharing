package main

import (
	"fmt"
	"reflect"

	"github.com/fairyhunter13/reflecthelper/v3"
)

// Get the elem type of a variable of reflect.Value
func main() {
	normal()
	usingReflectHelper()
}

func normal() {
	// normally
	var testSlice *[]int
	test := []int{1, 2, 3, 4, 5}
	testSlice = &test
	typeTestSlice := reflect.ValueOf(&testSlice).Type().Elem().Elem().Elem()
	newArr := reflect.New(reflect.ArrayOf(3, typeTestSlice)).Elem()
	newArr.Index(0).SetInt(3)
	fmt.Println("The new array is: ", newArr.Interface())
}

func usingReflectHelper() {
	// using reflecthelper
	var testSlice *[]int
	test := []int{1, 2, 3, 4, 5}
	testSlice = &test
	newArr := reflect.New(reflect.ArrayOf(
		3,
		reflecthelper.GetChildElemType(reflect.ValueOf(&testSlice)))).Elem()
	newArr.Index(0).SetInt(3)
	fmt.Println("The new array is: ", newArr.Interface())
}
