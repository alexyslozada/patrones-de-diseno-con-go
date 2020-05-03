package message

import "github.com/alexyslozada/patrones-de-diseno-con-go/comportamiento/observer/observer"

// Message .
type Message struct {
	observers map[string]observer.Observer
	Msg       string
}

// AddObserver .
func (m *Message) AddObserver(name string, o observer.Observer) {
	if m.observers == nil {
		m.observers = make(map[string]observer.Observer)
	}

	m.observers[name] = o
}

// RemoveObserver .
func (m *Message) RemoveObserver(name string) {
	delete(m.observers, name)
}

// NotifyObservers .
func (m *Message) NotifyObservers() {
	for _, v := range m.observers {
		v.Notify(m.Msg)
	}
}
