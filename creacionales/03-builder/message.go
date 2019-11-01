package builder

// Producto concreto
type Message struct {
	Body   []byte
	Format string
}

// Producto Base
type MessageFormat struct {
	Recipient   string `json:"recipient",xml:"recipient"`
	MessageText string `json:"message",xml:"message"`
}
