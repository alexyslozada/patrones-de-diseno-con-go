package remote

import (
	"time"

	"github.com/alexyslozada/pruebas/proxy/book"
)

// Data simula la base de datos
type Data struct {
	books    book.Books
	server   string
	port     uint16
	user     string
	password string
}

// New devuelve una nueva BD
func New(server string, port uint16, user, password string) *Data {
	d := &Data{
		server:   server,
		port:     port,
		user:     user,
		password: password,
	}
	d.load()
	return d
}

// ByID retorna un libro de la lista
func (d *Data) ByID(ID uint) book.Book {
	time.Sleep(2 * time.Second)
	for _, v := range d.books {
		if v.ID == ID {
			return v
		}
	}

	return book.Book{}
}

// All retorna todos los libros
func (d *Data) All() book.Books {
	time.Sleep(4 * time.Second)
	return d.books
}

// load Carga la data inicial
func (d *Data) load() {
	d.books = make(book.Books, 0, 5)
	d.books = append(
		d.books,
		book.Book{ID: 1, Name: "El arte de la guerra", Author: "Sun Tzu"},
		book.Book{ID: 2, Name: "La pelota no entra por azar", Author: "Ferran Soriano"},
		book.Book{ID: 3, Name: "Jesus, CEO", Author: "Laurie Beth"},
		book.Book{ID: 4, Name: "La biografía de Steve Jobs", Author: "Walter Isaacson"},
		book.Book{ID: 5, Name: "Pequeño cerdo capitalista", Author: "Sofía Macías"},
	)
}
