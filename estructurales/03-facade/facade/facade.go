package facade

import (
	"github.com/alexyslozada/patrones-de-diseno-con-go/estructurales/03-facade/email"
	"github.com/alexyslozada/patrones-de-diseno-con-go/estructurales/03-facade/storage"
	"github.com/alexyslozada/patrones-de-diseno-con-go/estructurales/03-facade/validate"
)

type Facade struct {
	to                  string
	comment             string
	validatorToken      validate.ValidatorToken
	validatorPermission validate.ValidatorPermission
	store               storage.Storage
	notificator         email.Email
}

func New(to, comment, token, user string) Facade {
	return Facade{
		to:                  to,
		comment:             comment,
		validatorToken:      validate.NewValidatorToken(token),
		validatorPermission: validate.NewValidatorPermission(user),
		store:               storage.NewStorage("mysql"),
		notificator:         email.NewEmail(),
	}
}

func (f *Facade) Comment() error {
	if err := f.validatorToken.Validate(); err != nil {
		return err
	}

	if err := f.validatorPermission.Validate(); err != nil {
		return err
	}

	f.store.Save(f.comment)

	f.notificator.Send(f.to, f.comment)

	return nil
}

func (f *Facade) Notify() {
	f.notificator.Send(f.to, f.comment)
}
