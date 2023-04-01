package state

import (
	"fmt"
	"os"
	"strings"
)

type PendingApprove struct{}

func (pa PendingApprove) Run(po *PurchaseOrder) bool {
	accepted := ""
	fmt.Println("Aceptas la orden? S/N")
	_, err := fmt.Fscanf(os.Stdin, "%s\n", &accepted)
	if err != nil {
		fmt.Println("no se pudo procesar la opción")
		return false
	}
	if !strings.EqualFold(accepted, "s") {
		po.State = Rejected{}
		return true
	}

	po.State = Accepted{}
	return true
}
