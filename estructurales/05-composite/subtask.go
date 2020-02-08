package main

import (
	"fmt"
	"strings"
)

// SubTask ...
type SubTask struct {
	name  string
	price int
}

// Add ...
func (s *SubTask) Add(i Item) {
	fmt.Println("Yo no acepto más subtareas")
}

// String ...
func (s *SubTask) String() string {
	sb := strings.Builder{}
	sb.WriteString("\t\t|-- ")
	sb.WriteString(s.name)
	sb.WriteString("\n")

	return sb.String()
}

// Price ...
func (s *SubTask) Price() int {
	return s.price
}
