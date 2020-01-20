package validate

import (
	"errors"
	"fmt"
)

var ErrTokenNotValid = errors.New("el usuario no está logueado")

type ValidatorToken struct {
	token string
}

func NewValidatorToken(t string) ValidatorToken {
	return ValidatorToken{token: t}
}

func (vt *ValidatorToken) Validate() error {
	if vt.token != "token-valido" {
		return ErrTokenNotValid
	}

	fmt.Println("token válido")
	return nil
}
