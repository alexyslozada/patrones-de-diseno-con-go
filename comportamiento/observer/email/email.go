package email

import "fmt"

// Email .
type Email struct {}

// Notify .
func (e *Email) Notify(data string) {
	sendEmail(data)
}

func sendEmail(data string) {
	fmt.Printf("He enviado un email con el mensaje: %q\n", data)
}
