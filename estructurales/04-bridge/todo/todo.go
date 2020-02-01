package todo

import "github.com/alexyslozada/patrones-de-diseno-con-go/estructurales/04-bridge/list"

// ToDo interface que tiene la lista de tareas
type ToDo interface {
	SetList(l list.List)
	Add(name string)
	Print()
}
