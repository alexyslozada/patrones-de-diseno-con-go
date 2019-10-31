package factory

import "github.com/alexyslozada/patrones-de-diseno-con-go/creacionales/02-factory/connection"

func Factory(t int) connection.DBConnection {
	switch t {
	case 1:
		return &connection.MySql{}
	case 2:
		return &connection.Postgres{}
	default:
		return nil
	}
}
