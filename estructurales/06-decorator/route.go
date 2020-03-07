package decorator

import "fmt"

type Route struct {
	decorators map[string]Decorator
}

func NewRoute() Route {
	return Route{make(map[string]Decorator)}
}

func (r *Route) Add(decorator Decorator, path string) {
	r.decorators[path] = decorator
}

func (r *Route) Exec(path string) {
	d, ok := r.decorators[path]
	if !ok {
		fmt.Println("no existe esa ruta")
		return
	}

	err := d.Process()
	if err != nil {
		fmt.Println("ERROR:", err)
	}
}
