package chain

import "fmt"

type College struct {
	supervisor Responsable
}

func NewCollege(s Responsable) College {
	return College{supervisor: s}
}

func (c College) Handle(process string) {
	if process == "CR" {
		fmt.Println("Revisando... Qué bien ✅")
		return
	}

	c.supervisor.Handle(process)
}
