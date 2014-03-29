package level

import "container/list"

import "github.com/jameseb7/roguelike/entity"
import "github.com/jameseb7/roguelike/symbol"

type cellType struct{
	baseSymbol symbol.Symbol
	occupant entity.ID
	items list.List
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