package entity

import "github.com/jameseb7/roguelike/symbol"

type ID uint64

type ConstEntity interface {
	EntityID() ID
	EntityName() string
	EntitySymbol() symbol.Symbol
	EntityItemType() ItemType
	Size() int
}

type Entity interface {
	ConstEntity
}

const (
	Small = iota
	Large
)

var currentID ID = 0

func NewEntityID() ID {
	currentID++
	if currentID == 0 {
		panic("Out of Entity IDs")
	}
	return currentID
}

type Context interface {
	EntityByID(eid ID) ConstEntity
}

type Actor interface {
	Entity
	NextAction(c Context) Action
}
