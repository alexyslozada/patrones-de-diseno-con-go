package main

import (
	"github.com/alexyslozada/patrones-de-diseno-con-go/comportamiento/state"
	"time"
)

func main() {
	initialized := state.Initialized{}
	po := state.PurchaseOrder{
		Client: "Alexys Lozada",
		Date:   time.Now(),
		State:  initialized,
	}

	for po.State.Run(&po) {
	}
}
