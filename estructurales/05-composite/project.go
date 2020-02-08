package main

import (
	"strconv"
	"strings"
)

// Project ...
type Project struct {
	name  string
	tasks []Item
}

// Add ...
func (p *Project) Add(i Item) {
	p.tasks = append(p.tasks, i)
}

// String ...
func (p *Project) String() string {
	sb := strings.Builder{}
	sb.WriteString(p.name)
	sb.WriteString(" $")
	sb.WriteString(strconv.Itoa(p.Price()))
	sb.WriteString("\n")
	for _, v := range p.tasks {
		sb.WriteString(v.String())
	}

	return sb.String()
}

// Price ...
func (p *Project) Price() int {
	price := 0
	for _, v := range p.tasks {
		price += v.Price()
	}

	return price
}
