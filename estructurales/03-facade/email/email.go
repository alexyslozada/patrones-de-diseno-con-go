package email

import "fmt"

type Email struct {
	to      string
	message string
}

func NewEmail() Email {
	return Email{}
}

func (e *Email) Send(to, comment string) {
	e.to = to
	e.message = comment
	fmt.Printf("Enviando mail a: %s, mensaje: %s\n", e.to, e.message)
}
