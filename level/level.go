package level

import "container/list"

import "github.com/jameseb7/roguelike/entity"
import "github.com/jameseb7/roguelike/action"
import "github.com/jameseb7/roguelike/symbol"
import "github.com/jameseb7/roguelike/direction"

type Level interface{
	SymbolAt(x, y int) symbol.Symbol

	Put(e entity.Entity, x, y int) (ok bool)
	Remove(eid entity.ID) (ok bool, e entity.Entity)

	Run() action.Action
}

const XWidth = 80
const YWidth = 20
type entityMetadata struct{
	xPosition int
	yPosition int
	entity entity.Entity
	turnSlot *list.Element
}

type baseLevel struct{
	cells [XWidth][YWidth]cellType
	entities map[entity.ID] *entityMetadata
	actors *list.List
	currentActor *list.Element
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

	if bl.cells[x][y].occupant != 0 {
		return false
	}

	bl.cells[x][y].occupant = e.EntityID()

	metadata := new(entityMetadata)
	metadata.xPosition = x
	metadata.yPosition = y
	metadata.entity = e

	if _, ok := e.(entity.Actor); ok {
		ts := bl.actors.PushBack(e.EntityID())
		metadata.turnSlot = ts
	}
	
	bl.entities[e.EntityID()] = metadata
	return true
}

func (bl *baseLevel) Remove(eid entity.ID) (ok bool, e entity.Entity) {
	metadata := bl.entities[eid]
	if metadata == nil {
		return false, nil
	}

	bl.cells[metadata.xPosition][metadata.yPosition].occupant = 0
	if metadata.turnSlot != nil {
		_ = bl.actors.Remove(metadata.turnSlot)
	}
	e = metadata.entity
	delete(bl.entities, eid)

	ok = true
	return
}

func (bl *baseLevel) Run() action.Action {
	for {
		for ; bl.currentActor != nil; bl.currentActor = bl.currentActor.Next() {
			eid := bl.currentActor.Value.(entity.ID)
			e := bl.entities[eid]
			if e == nil {
				bl.actors.Remove(bl.currentActor)
			}
			a := e.entity.(entity.Actor)
			switch act := a.NextAction().(type) {
			case action.Player: 
				return act
			case action.Move: 
				bl.move(eid, act.Dir)
			}
		}
		bl.currentActor = bl.actors.Front()
		if bl.currentActor == nil {
			return action.Skip{}
		}
	}
}

func (bl *baseLevel) move(eid entity.ID, dir direction.Direction) {
	switch dir {
	case direction.North, direction.South, direction.East, direction.West:
		fallthrough
	case direction.NorthEast, direction.NorthWest:
		fallthrough
	case direction.SouthEast, direction.SouthWest:
		metadata := bl.entities[eid]
		offset := direction.Directions[dir]
		newx := metadata.xPosition + offset.X
		newy := metadata.yPosition + offset.Y
		if (newx < 0) || (newx >= XWidth) || 
			(newy < 0) || (newy >= YWidth) {
			return
		}

		if bl.cells[newx][newy].blocksMovement() {
			return
		}

		
		bl.cells[newx][newy].occupant = eid
		bl.cells[metadata.xPosition][metadata.yPosition].occupant = 0
		metadata.xPosition = newx
		metadata.yPosition = newy
	}
	return
}