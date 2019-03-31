package iterator

type Aggregate interface {
	Iterator() Iterator
}

type Iterator interface {
	HasNext() bool
	Next() interface{}
}
