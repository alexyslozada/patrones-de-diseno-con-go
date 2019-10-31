package client_two

import "github.com/alexyslozada/patrones-de-diseno-con-go/creacionales/01-singleton/singleton"

func IncrementAge() {
	p := singleton.GetInstance()
	p.IncrementAge()
}