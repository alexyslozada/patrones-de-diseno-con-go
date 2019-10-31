package builder

type Sender struct {
	builder MessageBuilder
}

// SetBuilder asigna el constructor
func (s *Sender) SetBuilder(b MessageBuilder) {
	s.builder = b
}

// Build a concrete message via MessageBuilder
func (s *Sender) BuildMessage() (*Message, error) {
	s.builder.
		SetRecipient("Mi mejor amigo/a").
		SetMessage("Felicidades por aprender patrones de diseño en Go.")
	return s.builder.Build()
}
