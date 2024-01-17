package main

import "fmt"

/*
Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
*/

type Human struct {
	name string
	age  int
}

func (h *Human) GetName() string {
	return h.name
}

func (h *Human) GetAge() int {
	return h.age
}

type Action struct {
	// Непосредственно, применяем механизм встраивания
	Human
}

func main() {
	action := &Action{}
	fmt.Printf("Age: %d\n", action.GetAge())
}
