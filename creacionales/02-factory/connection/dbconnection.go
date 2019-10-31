package connection

import "time"

type DBConnection interface {
	Connect() error
	GetNow() (*time.Time, error)
	Close() error
}
