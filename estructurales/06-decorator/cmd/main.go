package main

import (
	"fmt"
	"os"

	"github.com/alexyslozada/pruebas/decorator"
)

func main() {
	route := decorator.NewRoute()
	start(&route)

	var path string
	_, err := fmt.Scanf("%s", &path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	route.Exec(path)
}

func start(route *decorator.Route) {
	route.Add(decorator.NewHandlerSayHello(), "/saludar")
	route.Add(decorator.NewHandlerSayGoodbye(), "/despedirse")
}
