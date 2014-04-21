package level

import "container/list"
import "log"

import "github.com/jameseb7/roguelike/entity"
import "github.com/jameseb7/roguelike/symbol"
import "github.com/jameseb7/roguelike/direction"

type Level interface{
	entity.Context
	
	SymbolAt(x, y int) symbol.Symbol
	ItemsAt(x, y int) []entity.ID

	EntityLocation(eid entity.ID) (x, y int) 

	Put(e entity.Entity, x, y int) (ok bool)
	Remove(eid entity.ID) (ok bool, e entity.Entity)

	Run() entity.Action
	
	Turn() int
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

	turn int
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

	itemElement := bl.cells[x][y].items.Front()
	if itemElement != nil {
		itemID := itemElement.Value.(entity.ID)
		if bl.entities[itemID] != nil {
			if bl.entities[itemID].entity != nil {
				return bl.entities[itemID].entity.EntitySymbol()
			}
		}
	}
	
	return bl.cells[x][y].baseSymbol
}

func (bl *baseLevel) ItemsAt(x,y int) []entity.ID {
	if (x < 0) || (x >= XWidth) || (y < 0) || (y >= YWidth) {
		return nil
	}

	itemSlice := make([]entity.ID, 0, bl.cells[x][y].items.Len())
	for item := bl.cells[x][y].items.Front(); item != nil; item = item.Next() {
		if eid := item.Value.(entity.ID); bl.entities[eid] != nil {
			itemSlice = append(itemSlice, eid)
		} else {
			//nonexistent items shouldn't be in the cell item list
			tmp := item.Prev()
			bl.cells[x][y].items.Remove(item)
			item = tmp
		}
	}

	return itemSlice
}

func (bl *baseLevel) Put(e entity.Entity, x, y int) (ok bool) {
	if (x < 0) || (x >= XWidth) || (y < 0) || (y >= YWidth) {
		return false
	}
	if e.Size() >= entity.Large {
		if bl.cells[x][y].occupant != 0 {
			return false
		}
		
		bl.cells[x][y].occupant = e.EntityID()
	} else {
		bl.cells[x][y].items.PushFront(e.EntityID())
	}

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

func (bl *baseLevel) Run() entity.Action {
	for {
		if bl.currentActor == nil {
			bl.currentActor = bl.actors.Front()
			if bl.currentActor == nil {
				return entity.SkipAction{}
			}
		}
		for ; bl.currentActor != nil; bl.currentActor = bl.currentActor.Next() {
			eid := bl.currentActor.Value.(entity.ID)
			e := bl.entities[eid]
			if e == nil {
				bl.actors.Remove(bl.currentActor)
			}
			a := e.entity.(entity.Actor)
			for actionDone := false; !actionDone; {
				switch act := a.NextAction(bl).(type) {
				case entity.PlayerAction: 
					return act
				case entity.MoveAction: 
					unresolved, impossible := bl.move(eid, act.Dir)
					if unresolved {
						return act
					}
					if !impossible {
						actionDone = true
					}
				case entity.SkipAction:
					actionDone = true
				case entity.PickUpAction:
					x, y := e.xPosition, e.yPosition
					for _, v := range act.Items {
						if ok := bl.cells[x][y].removeEntity(v); ok { //check the requested entity is available to pickup
							c := a.(entity.Container)
							if item := bl.entities[v].entity; item != nil { //check that the requested entity actually exists
								if stackable, ok := item.(entity.Stackable); ok {
									//handle stackable items
									log.Println("Item", v, "is stackable")
									stackEID := c.ItemOfType(stackable.EntityItemType())
									if bl.entities[stackEID] != nil { 
										log.Println("Already a stackable item", stackEID, "of the same type as", v)
										stackEntity := bl.entities[stackEID].entity.(entity.Stackable)
										if stackEntity.EntityItemType() == stackable.EntityItemType() {
											stackEntity.Add(stackable.Count())
											log.Println("Total items in stack:", stackEntity.Count()) 
										}
									} else {
										c.AddItem(v, item.EntityItemType())
									}
								} else {
									c.AddItem(v, item.EntityItemType())
								}
							}
						}
					}
				default:
					return act
				}
			}
			bl.turn++
		}
	}
}

func (bl *baseLevel) Turn() int {
	return bl.turn
}

func (bl *baseLevel) move(eid entity.ID, dir direction.Direction) (unresolved, impossible bool){
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
		if (newx < 0) || (newx >= XWidth) || (newy < 0) || (newy >= YWidth) {
			impossible = true
			return
		}

		if bl.cells[newx][newy].blocksMovement() {
			impossible = true
			return
		}

		
		bl.cells[newx][newy].occupant = eid
		bl.cells[metadata.xPosition][metadata.yPosition].occupant = 0
		metadata.xPosition = newx
		metadata.yPosition = newy
	default:
		unresolved = true
		return
	}
	return
}

func (bl *baseLevel) EntityByID(eid entity.ID) entity.ConstEntity {
	return bl.entities[eid].entity
}

func (bl *baseLevel) EntityLocation(eid entity.ID) (x, y int) {
	e := bl.entities[eid]
	if e != nil {
		return e.xPosition, e.yPosition
	}
	return -1, -1
}