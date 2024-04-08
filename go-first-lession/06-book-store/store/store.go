package store

import "errors"

var (
	ErrNotFound = errors.New("not found")
	ErrExist    = errors.New("exist")
)

type Book struct {
	Id      string   `json:"id"`      // ISBN ID
	Name    string   `json:"name"`    // name
	Authors []string `json:"authors"` // authors
	Press   string   `json:"press"`   // press
}

type Store interface {
	Create(*Book) error //
	Update(*Book) error
	Get(string) (Book, error)
	GetAll() ([]Book, error)
	Delete(string) error
}
