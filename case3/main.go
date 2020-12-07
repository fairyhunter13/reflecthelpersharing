package main

import (
	"fmt"
	"reflect"

	"github.com/fairyhunter13/reflecthelper/v3"
)

// Get the child elem kind of a multi level pointer variable of reflect.Value
func main() {
	normal()
	usingReflectHelper()
}

func normal() {
	// normally
	var test ***int
	kindOfTest := reflect.ValueOf(&test).Type().Elem().Elem().Elem().Elem().Kind()
	fmt.Println("The child elem kind of test is:", kindOfTest)
}

func usingReflectHelper() {
	// using reflecthelper
	var test ***int
	kindOfTest := reflecthelper.GetChildElemPtrKind(reflect.ValueOf(&test))
	fmt.Println("The child elem kind of test is:", kindOfTest)
}
