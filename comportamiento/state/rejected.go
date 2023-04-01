package state

import (
	"fmt"
	"os"
	"strings"
)

type Rejected struct{}

func (r Rejected) Run(po *PurchaseOrder) bool {
	rejected := ""
	fmt.Println("Has rechazado la orden, estás seguro de que quieres rechazarla? S/N")
	_, err := fmt.Fscanf(os.Stdin, "%s\n", &rejected)
	if err != nil {
		fmt.Println("no se pudo procesar la opción")
		return false
	}
	if !strings.EqualFold(rejected, "s") {
		po.State = PendingApprove{}
		return true
	}

	fmt.Printf("%s, Lamentamos que hayas rechazado la oferta 😢\n", po.Client)
	return false
}
