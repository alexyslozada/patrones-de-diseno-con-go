package local

import (
	"github.com/alexyslozada/pruebas/proxy/book"
	"github.com/alexyslozada/pruebas/proxy/remote"
)

type Local struct {
	Remote *remote.Data
	cache  book.Books
}

func New() Proxy {
	return &Local{
		Remote: remote.New("https://alguno.com", 8080, "usuario", "123456"),
		cache:  make(book.Books, 0),
	}
}

func (l *Local) GetByID(ID uint) book.Book {
	for _, v := range l.cache {
		if v.ID == ID {
			return v
		}
	}

	b := l.Remote.ByID(ID)
	l.cache = append(l.cache, b)

	return b
}

func (l *Local) GetAll() book.Books {
	return l.Remote.All()
}
