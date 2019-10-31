package builder

type MessageBuilder interface {
	SetRecipient(recipient string) MessageBuilder
	SetMessage(message string) MessageBuilder
	Build() (*Message, error)
}
