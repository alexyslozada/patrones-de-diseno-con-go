package slack

import "testing"

func Test_SendMessage(t *testing.T) {
	data := "Hola desde test de go"
	sendMessage(data)
}
