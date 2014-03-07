package level

import "container/list"

import "entity"
import "action"
import "symbol"

type Level interface{
	SymbolAt(x, y int) symbol.Symbol

	Put(e entity.Entity, x, y int) (ok bool)
	Remove(e entity.ID) (ok bool)

	Run() action.Action
}

const XWidth = 80
const YWidth = 20

type cellType struct{
	baseSymbol symbol.Symbol
	occupant entity.ID
}

type entityMetadata struct{
	xPosition int
	yPosition int
	entity entity.Entity
	turnSlot *list.Element
}

type baseLevel struct{
	cells [XWidth][YWidth]cellType
	entities map[entity.ID] *entityMetadata
	actors list.List
}

func (bl *baseLevel) SymbolAt(x,y int) symbol.Symbol {
	if (x < 0) || (x >= XWidth) || (y < 0) || (y >= YWidth) {
		return symbol.Blank
	}

	occupantID := bl.cells[x][y].occupant
	if bl.entities[occupantID] != nil {
		if bl.entities[occupantID].entity != nil {
			return bl.entities[occupantID].entity.EntitySymbol()
		}
	}
	return bl.cells[x][y].baseSymbol
}

func (bl *baseLevel) Put(e entity.Entity, x, y int) (ok bool) {
	if (x < 0) || (x >= XWidth) || (y < 0) || (y >= YWidth) {
		return false
	}

	if bl.cells[x][y].occupant != null {
		return false
	}

	bl.cells[x][y].occupant = e.EntityID()

	metadata := new(entityMetadata)
	metadata.xPosition = x
	metadata.yPosition = y
	metadata.entity = e

	if _, ok := e.(Actor); ok {
		ts := bl.actors.PushBack(e.EntityID())
		metadata.turnSlot = ts
	}
	
	bl.entities[e.EntityID()] = metadata
	return true
}