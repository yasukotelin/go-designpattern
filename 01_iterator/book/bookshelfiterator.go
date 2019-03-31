package book

type BookShelfIterator struct {
	bookShelf *BookShelf
	index     int
}

func NewBookShelfIterator(shelf *BookShelf) *BookShelfIterator {
	return &BookShelfIterator{
		bookShelf: shelf,
		index:     0,
	}
}

func (iterator *BookShelfIterator) HasNext() bool {
	if iterator.index < iterator.bookShelf.GetLength() {
		return true
	}
	return false
}

func (iterator *BookShelfIterator) Next() interface{} {
	book := iterator.bookShelf.GetBookAt(iterator.index)
	iterator.index++
	return book
}
