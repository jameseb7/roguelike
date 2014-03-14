package entity

import "github.com/jameseb7/roguelike/action"
import "github.com/jameseb7/roguelike/symbol"

type ID uint64

type Entity interface {
	EntityID() ID
	EntitySymbol() symbol.Symbol
}

var currentID ID = 0

func NewEntityID() ID {
	currentID++
	if currentID == 0 {
		panic("Out of Entity IDs")
	}
	return currentID
}

type Actor interface {
	Entity
	NextAction() action.Action
}
