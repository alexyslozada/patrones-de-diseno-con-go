package state

type State interface {
	Run(ctx *PurchaseOrder) bool
}
