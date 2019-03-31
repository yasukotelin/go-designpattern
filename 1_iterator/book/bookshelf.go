package book

import "github.com/yasukotelin/go-designpattern/1_iterator/iterator"

type BookShelf struct {
	books []Book
	last  int
}

func NewBookShelf(max int) *BookShelf {
	return &BookShelf{
		books: make([]Book, max),
	}
}

func (shelf *BookShelf) GetBookAt(index int) *Book {
	return &shelf.books[index]
}

func (shelf *BookShelf) AppendBook(book *Book) {
	shelf.books[shelf.last] = *book
	shelf.last++
}

func (shelf *BookShelf) GetLength() int {
	return shelf.last
}

func (shelf *BookShelf) Iterator() iterator.Iterator {
	return NewBookShelfIterator(shelf)
}
