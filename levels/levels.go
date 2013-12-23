package levels

import "github.com/jameseb7/roguelike/types"

type LevelType int

const (
	TEST = iota
)

func Make(t LevelType, parent types.Region, index int) types.Level {
	switch t {
	case TEST:
		var b = new(baseLevel)
		for x := 0; x < b.XWidth(); x++ {
			for y := 0; y < b.YWidth(); y++ {
				if x == 0 || x == b.XWidth()-1 {
					b.cells[x][y].cellType = types.VWALL
				} else if y == 0 || y == b.YWidth()-1 {
					b.cells[x][y].cellType = types.HWALL
				} else {
					b.cells[x][y].cellType = types.FLOOR
				}
			}
		}
		b.parent = parent
		b.index = index
		return b
	default:
		return nil
	}
}
