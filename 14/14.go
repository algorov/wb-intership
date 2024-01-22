package main

import (
	"fmt"
	"reflect"
)

// Разработать программу, которая в рантайме способна определить
// тип переменной: int, string, bool, channel из переменной типа interface{}.

func main() {
	// Для демонстрации инициализируется последовательность элементов разного типа.
	elems := []interface{}{1, "delightful", true, make(chan int)}

	// Определение типа каждого элемента.
	for _, elem := range elems {
		fmt.Println(Reflection(elem))
	}
}

// Определение посредством рефлексии.
func Reflection(elem interface{}) string {
	switch reflect.TypeOf(elem).String() {
	case "int":
		return "int"
	case "string":
		return "string"
	case "bool":
		return "bool"
	case "chan int":
		return "channel"
	}

	return "undefined"
}
