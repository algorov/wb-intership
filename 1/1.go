package main

import "fmt"

/*
Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
*/

// Вот такая простенькая структура.
type Human struct {
	name string
	age  int
}

// Метод, который привязан (ссылается) к структуре Human.
func (h *Human) GetName() string {
	return h.name
}

// Метод, который привязан (ссылается) к структуре Human.
func (h *Human) GetAge() int {
	return h.age
}

type Action struct {
	// Непосредственно, применяем механизм встраивания.
	Human
}

func main() {
	action := &Action{}
	fmt.Printf("Age: %d\n", action.GetAge())
}
