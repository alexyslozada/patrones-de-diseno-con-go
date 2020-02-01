package list

import "fmt"

type Bullet struct {
	bullet rune
}

// Retorna una nueva representación de listas
func NewBullet(b rune) *Bullet {
	return &Bullet{
		bullet: b,
	}
}

func (b *Bullet) Print(todos []string) {
	for _, v := range todos {
		fmt.Println("\t", string(b.bullet), v)
	}
}
