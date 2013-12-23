package levels

import "github.com/jameseb7/roguelike/types"

const defaultXWidth = 80
const defaultYWidth = 20

type cell struct {
	cellType types.Symbol
	occupant types.Entity
}

type baseLevel [defaultXWidth][defaultYWidth]cell

func (bl baseLevel) SymbolAt(x, y int) types.Symbol {
	if bl[x][y].occupant != nil {
		return bl[x][y].occupant.Symbol()
	}

	if x < 0 || x > defaultXWidth {
		return types.BLANK
	}
	if y < 0 || y > defaultYWidth {
		return types.BLANK
	}

	return bl[x][y].cellType
}

func (bl baseLevel) XWidth() int { return defaultXWidth }
func (bl baseLevel) YWidth() int { return defaultYWidth }

func (bl baseLevel) IsOccupied(x, y int) bool {
	return bl[x][y].occupant != nil
}

func (bl *baseLevel) Put(e types.Entity, x, y int) (ok bool) {
	if x < 0 || x >= defaultXWidth || y < 0 || y >= defaultYWidth {
		return false
	}

	if bl.IsOccupied(x, y) {
		return false
	}

	bl[x][y].occupant = e
	e.SetX(x)
	e.SetY(y)
	e.SetParent(bl)
	return true
}

func (bl *baseLevel) Move(e types.Entity, dir types.Direction) (ok bool) {
	x, y := e.X(), e.Y()

	//check the entity can sucessfully be placed in the new location first
	dirXYZ := types.Directions[dir]
	ok = bl.Put(e, x+dirXYZ.X, y+dirXYZ.Y)
	if ok == false {
		return
	}

	//now remove the entity from its old location
	bl[x][y].occupant = nil

	return
}
