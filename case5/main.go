package main

import (
	"fmt"
	"reflect"

	"github.com/fairyhunter13/reflecthelper/v3"
)

// Assignment of two variables of reflect.Value with different type
func main() {
	normal()
	usingReflectHelper()
}

func normal() {
	// normally
	uintVal := uint64(65)
	var intVal *int
	hello := 0
	intVal = &hello
	assignerRef := reflect.ValueOf(&intVal).Elem().Elem()
	valRef := reflect.ValueOf(uintVal)
	assignerRef.SetInt(int64(valRef.Uint()))
	fmt.Println("The intVal is:", *intVal)
}

func usingReflectHelper() {
	// using reflecthelper
	uintVal := uint64(65)
	var intVal *int
	err := reflecthelper.AssignReflect(
		reflecthelper.GetInitChildElem(reflect.ValueOf(&intVal)),
		reflect.ValueOf(uintVal),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("The intVal is:", *intVal)
}
