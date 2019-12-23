package main

import (
	"fmt"
	"time"

	"github.com/alexyslozada/pruebas/proxy/local"
)

var loc local.Proxy

func main() {
	loc = local.New()

	GetByID(2)
	GetByID(2)
	GetByID(1)
	GetByID(2)
	GetByID(3)
	GetByID(2)
	GetByID(1)
	GetAll()
}

// GetByID permite ver un libro por su ID
func GetByID(ID uint) {
	start := time.Now()
	fmt.Printf("%+v", loc.GetByID(ID))
	elapsed := time.Since(start)
	fmt.Printf("Tiempo usado: %v\n", elapsed)
}

// GetAll obtiene todos los libros
func GetAll() {
	start := time.Now()
	fmt.Printf("%+v", loc.GetAll())
	elapsed := time.Since(start)
	fmt.Printf("Tiempo usado: %v\n", elapsed)
}
