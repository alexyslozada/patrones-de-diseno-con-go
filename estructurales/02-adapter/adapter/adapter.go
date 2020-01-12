package adapter

import (
	"github.com/alexyslozada/patrones-de-diseno-con-go/estructurales/02-adapter/auto"
	"github.com/alexyslozada/patrones-de-diseno-con-go/estructurales/02-adapter/bici"
)

type Adapter interface {
	Mover()
}

type AutomovilAdapter struct {
	Auto *auto.Automovil
}

func (a *AutomovilAdapter) Mover() {
	if !a.Auto.Encendido {
		a.Auto.Encender()
	}

	a.Auto.Acelerar()
}

type BicicletaAdapter struct {
	Bici *bici.Bicicleta
}

func (b *BicicletaAdapter) Mover() {
	b.Bici.Avanzar()
}
