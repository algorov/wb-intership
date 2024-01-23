package main

// Реализовать паттерн «адаптер» на любом примере.

import "fmt"

type OldService struct {
	Username string
}

func (l *OldService) Operation() string {
	return fmt.Sprintf("Operation performed by %s", l.Username)
}

// NewService интерфейс, который мы хотим использовать в клиентском коде.
type NewService interface {
	NewOperation() string
}

// Adapter - адаптер, который адаптирует OldService под интерфейс NewService.
type Adapter struct {
	old *OldService
}

func (a *Adapter) NewOperation() string {
	// Преобразовывает результат из старого сервиса в формат нового.
	return "New operation performed by " + a.old.Operation()
}

// Работа, непосредственно, с новым сервисом, который адаптирован под старый.
func ClientCode(service NewService) {
	fmt.Println(service.NewOperation())
}

func main() {
	// Экземпляр OldService.
	legacyService := &OldService{Username: "Alice"}

	// Сам адаптер.
	adapter := &Adapter{old: legacyService}

	ClientCode(adapter)
}
