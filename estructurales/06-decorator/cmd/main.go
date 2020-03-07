package main

import (
	"fmt"
	"os"

	decorator "github.com/alexyslozada/patrones-de-diseno-con-go/estructurales/06-decorator"
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
	route.Add(decorator.NewLogRegistry(decorator.NewHandlerSayHello()), "/saludar")
	route.Add(decorator.NewLogRegistry(decorator.NewHandlerSayGoodbye()), "/despedirse")
}
