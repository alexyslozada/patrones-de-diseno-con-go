package builder_test

import (
	builder "github.com/alexyslozada/patrones-de-diseno-con-go/creacionales/03-builder"
	"testing"
)

func TestSender_BuildMessage(t *testing.T) {
	json := &builder.JSONMessageBuilder{}
	xml := &builder.XMLMessageBuilder{}
	sender := &builder.Sender{}

	sender.SetBuilder(json)
	jsonMsg, err := sender.BuildMessage("Gopher", "Hola mundo!")
	if err != nil {
		t.Fatalf("BuildMesage(): JSON: no se esperaba error, se recibió: %v", err)
	}

	t.Log(string(jsonMsg.Body))

	sender.SetBuilder(xml)
	xmlMsg, err := sender.BuildMessage("Gopher", "Hola mundo")
	if err != nil {
		t.Fatalf("BuildMesage(): XML: no se esperaba error, se recibió: %v", err)
	}

	t.Log(string(xmlMsg.Body))
}
