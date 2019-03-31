package main

import (
	"fmt"

	"github.com/yasukotelin/go-designpattern/1_iterator/book"
)

func main() {
	bookShelf := book.NewBookShelf(4)
	bookShelf.AppendBook(book.NewBook("Around the World in 80 Days"))
	bookShelf.AppendBook(book.NewBook("Bible"))
	bookShelf.AppendBook(book.NewBook("Cinderella"))
	bookShelf.AppendBook(book.NewBook("Daddy-Long-Legs"))

	iterator := bookShelf.Iterator()

	for iterator.HasNext() {
		if book, ok := iterator.Next().(*book.Book); ok {
			fmt.Println(book.Name)
		}
	}
}
