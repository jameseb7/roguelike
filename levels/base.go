package levels

import "github.com/jameseb7/roguelike/types"

const defaultXWidth = 80
const defaultYWidth = 20

type cell struct {
	cellType types.Symbol
	occupant types.Entity
}

func blocksPlacement(cellType types.Symbol) bool {
	switch cellType {
	case types.HWALL, types.VWALL:
		return true
	default:
		return false
	}
}

type baseLevel struct {
	index int
	parent types.Region
	cells [defaultXWidth][defaultYWidth]cell
	nextLevels [types.NUM_DIRECTIONS]types.Level
}

func (bl baseLevel) Index() int {
	return bl.index
}

func (bl baseLevel) SymbolAt(x, y int) types.Symbol {
	if bl.cells[x][y].occupant != nil {
		return bl.cells[x][y].occupant.Symbol()
	}

	if x < 0 || x > defaultXWidth {
		return types.BLANK
	}
	if y < 0 || y > defaultYWidth {
		return types.BLANK
	}

	return bl.cells[x][y].cellType
}

func (bl baseLevel) XWidth() int { return defaultXWidth }
func (bl baseLevel) YWidth() int { return defaultYWidth }

func (bl baseLevel) IsOccupied(x, y int) bool {
	return bl.cells[x][y].occupant != nil
}

func (bl *baseLevel) Put(e types.Entity, x, y int) (ok bool) {
	if x < 0 || x >= defaultXWidth || y < 0 || y >= defaultYWidth {
		return false
	}


	if bl.IsOccupied(x, y) {
		return false
	}
	if blocksPlacement(bl.cells[x][y].cellType) {
		return false
	}
	

	bl.cells[x][y].occupant = e
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
	bl.cells[x][y].occupant = nil

	return
}

func (bl *baseLevel) NextLevel(dir types.Direction) types.Level {
	if bl.nextLevels[dir] == nil {
		bl.nextLevels[dir] = bl.parent.NextLevel(bl.Index(), dir)
	}
	return bl.nextLevels[dir]
}
