package chain

import "fmt"

type SRE struct{}

func NewSRE() SRE {
	return SRE{}
}

func (sre SRE) Handle(process string) {
	if process == "DP" {
		fmt.Print("Deployando... 🚀")
		return
	}

	fmt.Println("No existe esa opción")
}
