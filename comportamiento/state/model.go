package state

import "time"

type PurchaseOrder struct {
	Client string
	Date   time.Time
	State  State
}
