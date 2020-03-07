package decorator

import "fmt"

type HandlerSayGoodbye struct{}

func NewHandlerSayGoodbye() *HandlerSayGoodbye {
	return &HandlerSayGoodbye{}
}

func (h *HandlerSayGoodbye) Process() error {
	fmt.Println("Adios")
	return nil
}
