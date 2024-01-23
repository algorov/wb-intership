package main

import (
	"firstLevel/24/point"
	"fmt"
)

func main() {
	// Экземпляры точек.
	fitstPoint := point.New(1, 5)
	secondPoint := point.New(0, 7)

	// Выводит расстояние между точками.
	fmt.Println(fitstPoint.Distance(secondPoint))
}
