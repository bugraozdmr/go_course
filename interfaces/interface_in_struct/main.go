package main

import "fmt"

// Interface
type Logger interface {
	Log(message string)
}

// ConsoleLogger, Logger interface'ini uygular
type ConsoleLogger struct{}

func (c ConsoleLogger) Log(message string) {
	fmt.Println("Log:", message)
}

// Service, Logger kullanır
type Service struct {
	logger Logger
}

func (s Service) DoSomething() {
	s.logger.Log("Bir işlem yapıldı.")
}

func main() {
	logger := ConsoleLogger{}
	service := Service{logger: logger}

	service.DoSomething() // Log: Bir işlem yapıldı.
}
