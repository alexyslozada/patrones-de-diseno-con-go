package main

import "fmt"

func main() {
	color := "rojo"
	phones := map[string]string{"home": "123456", "work": "789456"}
	p1 := prototype{"Alexys", 39, []string{"Deibis", "Carol"}, &color, phones}
	// copia del objeto original
	p2 := p1.Clone()

	// cambio de los valores al objeto p2
	p2.age = 40
	p2.name = "Carol"
	p2.friends[0] = "Maria"
	p2.friends[1] = "Pedro"

	// Estos cambios solo afectan al objeto p1
	color = "azul"
	phones["home"] = "147258"

	// muestra la información
	fmt.Println(p1.String())
	fmt.Println(p2.String())
}

