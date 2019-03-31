package book

type Book struct {
	Name string
}

func NewBook(name string) *Book {
	return &Book{
		Name: name,
	}
}
