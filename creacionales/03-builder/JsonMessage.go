package builder

import "encoding/json"

// JSON Message Builder is concrete builder
type JSONMessageBuilder struct {
	messageRecipient string
	messageText      string
}

func (b *JSONMessageBuilder) SetRecipient(recipient string) MessageBuilder {
	b.messageRecipient = recipient
	return b
}

func (b *JSONMessageBuilder) SetMessage(text string) MessageBuilder {
	b.messageText = text
	return b
}

func (b *JSONMessageBuilder) Build() (*Message, error) {
	m := MessageFormat{
		Recipient:   b.messageRecipient,
		MessageText: b.messageText,
	}

	data, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}

	return &Message{Body: data, Format: "JSON"}, nil
}
