package main

import (
	"fmt"
	"github.com/alexyslozada/patrones-de-diseno-con-go/comportamiento/chain"
	"log"
)

func main() {
	sre := chain.NewSRE()
	teamLead := chain.NewTeamLeader(sre)
	college := chain.NewCollege(teamLead)

	fmt.Println("Qué quieres solicitar? CR (code review), MR (merge request), DP (deploy)")
	option := ""
	_, err := fmt.Scanln(&option)
	if err != nil {
		log.Fatalf("cannot read your option: %v", err)
	}

	college.Handle(option)
}
