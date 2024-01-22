package main

import (
	"context"
	"fmt"
	"time"
)

// Реализовать собственную функцию sleep.

func main() {
	SleepWithContext(5)
	SleepWithTimer(5)
}

func SleepWithContext(duration int) {
	// Инициализирует контекст с таймером.
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(duration)*time.Second)
	fmt.Println("Run...")

	// Ждёт, пока не прочитается из закрытого канала.
	<-ctx.Done()
	fmt.Println("Done.")
}

func SleepWithTimer(duration int) {
	// Инициализирует таймер, который закрывает канал по истечению времени.
	timer := time.After(time.Duration(duration) * time.Second)
	fmt.Println("Run...")

	// Ждёт, пока не прочитается из закрытого канала.
	<-timer
	fmt.Println("Done.")
}
