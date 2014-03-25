package level

import "github.com/jameseb7/roguelike/entity"
import "github.com/jameseb7/roguelike/symbol"

type cellType struct{
	baseSymbol symbol.Symbol
	occupant entity.ID
}

func (c cellType) blocksMovement() bool {
	if c.occupant != 0 {
		return true
	}
	
	switch c.baseSymbol {
	case symbol.HWall, symbol.VWall:
		return true
	default:
		return false
	}
}