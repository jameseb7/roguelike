package entity

import "github.com/jameseb7/roguelike/action"
import "github.com/jameseb7/roguelike/symbol"

type ID uint64

type Entity interface {
	EntityID() ID
	EntityName() string
	EntitySymbol() symbol.Symbol

	Size() int
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
	EntityNameByID(eid ID) string
	EntitySymbolByID(eid ID) symbol.Symbol
	EntitySizeByID(eid ID) int
	EntityContents(eid ID) []ID
}

type Actor interface {
	Entity
	NextAction(c Context) action.Action
}
