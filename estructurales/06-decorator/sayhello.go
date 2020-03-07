package decorator

import "fmt"

type HandlerSayHello struct{}

func NewHandlerSayHello() *HandlerSayHello {
	return &HandlerSayHello{}
}

func (h *HandlerSayHello) Process() error {
	fmt.Println("Hola")
	return nil
}
