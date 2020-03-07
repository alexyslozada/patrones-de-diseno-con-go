package decorator

import "fmt"

type LogRegistry struct {
	Handler Decorator
}

func NewLogRegistry(handler Decorator) *LogRegistry {
	return &LogRegistry{handler}
}

func (lr *LogRegistry) Process() error {
	defer fmt.Println("Petición")
	return lr.Handler.Process()
}
