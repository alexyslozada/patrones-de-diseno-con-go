package list

import "fmt"

// Plain imprime el listado plano
type Plain struct{}

func NewPlain() *Plain {
	return &Plain{}
}

func (p *Plain) Print(todos []string) {
	for _, v := range todos {
		fmt.Println("\t", v)
	}
}
