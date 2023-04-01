package state

type Initialized struct{}

func (i Initialized) Run(po *PurchaseOrder) bool {
	po.State = PendingApprove{}
	return true
}
