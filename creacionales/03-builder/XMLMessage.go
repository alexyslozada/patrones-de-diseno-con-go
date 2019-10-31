package builder

import "encoding/xml"

// XML Message Builder is concrete builder
type XMLMessageBuilder struct {
	messageRecipient string
	messageText      string
}

func (b *XMLMessageBuilder) SetRecipient(recipient string) MessageBuilder {
	b.messageRecipient = recipient
	return b
}

func (b *XMLMessageBuilder) SetMessage(text string) MessageBuilder {
	b.messageText = text
	return b
}

func (b *XMLMessageBuilder) Build() (*Message, error) {
	m := MessageFormat{
		Recipient: b.messageRecipient,
		MessageText:      b.messageText,
	}

	data, err := xml.Marshal(m)
	if err != nil {
		return nil, err
	}

	return &Message{Body: data, Format: "XML"}, nil
}
