package main

import (
	"fmt"
	"log"

	"github.com/alexyslozada/patrones-de-diseno-con-go/estructurales/03-facade/facade"
)

func main() {

	showInit()

	token := "token-valido"
	user := "user-blog"
	to := "a@algo.com"
	comment := "muy buen video :)"

	f := facade.New(to, comment, token, user)
	err := f.Comment()
	if err != nil {
		log.Fatal(err)
	}

	f.Notify()

	showFinish()
}

func showInit() {
	fmt.Println(`
	**************************
	* Bienvenido al programa *
	**************************
	`)
}

func showFinish() {
	fmt.Println(`
	**************************
	*  Gracias por utilizar  *
	**************************
	`)
}
