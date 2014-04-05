package level

import "container/list"

import "github.com/jameseb7/roguelike/entity"
import "github.com/jameseb7/roguelike/symbol"

type cellType struct{
	baseSymbol symbol.Symbol
	occupant entity.ID
	items *list.List
}

func (c cellType) blocksMovement() bool {
	if c.occupant != 0 {
		return true
	}
	
	switch c.baseSymbol {
	case symbol.HWall, symbol.VWall, symbol.Rock:
		return true
	default:
		return false
	}
}

func (c cellType) removeEntity(eid entity.ID) (ok bool) {
	if c.occupant == eid {
		c.occupant = 0
		return true
	}

	for i := c.items.Front(); i != nil; i = i.Next() {
		if i.Value.(entity.ID) == eid {
			c.items.Remove(i)
			return true
		}
	}

	return false
} 