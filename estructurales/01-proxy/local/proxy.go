package local

import (
	"github.com/alexyslozada/pruebas/proxy/book"
)

type Proxy interface {
	GetByID(ID uint) book.Book
	GetAll() book.Books
}
