package main

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/fairyhunter13/reflecthelper/v3"
)

// Extract value with specific type from a variable of reflect.Value
func main() {
	normal()
	usingReflectHelper()
}

func normal() {
	// normally
	testVal := reflect.ValueOf("1")
	result, err := strconv.ParseBool(testVal.String())
	if err != nil {
		panic(err)
	}
	fmt.Println("The bool value of testVal:", result)
}

func usingReflectHelper() {
	// using reflecthelper
	result, err := reflecthelper.ExtractBool(reflect.ValueOf("false"))
	if err != nil {
		panic(err)
	}
	fmt.Println("The bool value of testVal:", result)
}
