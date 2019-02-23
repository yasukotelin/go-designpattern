package subject

import "github.com/yasukotelin/go-designpattern/observer/observer"

type Subject interface {
	Register(o observer.Observer)
	Remove(o observer.Observer)
	Notify()
}
