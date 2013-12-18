package levels

import "github.com/jameseb7/roguelike/types"

const defaultXWidth = 80
const defaultYWidth = 20

type cell struct {
	cellType types.Symbol
}

type baseLevel [defaultXWidth][defaultYWidth]cell

func (bl baseLevel) SymbolAt(x, y int) types.Symbol {
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
