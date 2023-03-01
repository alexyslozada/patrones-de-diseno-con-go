package chain

import "fmt"

type TeamLeader struct {
	supervisor Responsable
}

func NewTeamLeader(s Responsable) TeamLeader {
	return TeamLeader{supervisor: s}
}

func (tl TeamLeader) Handle(process string) {
	if process == "MR" {
		fmt.Print("Procesando merge... OK 🏘️")
		return
	}

	tl.supervisor.Handle(process)
}
