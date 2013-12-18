package levels

import "github.com/jameseb7/roguelike/types"

type LevelType int

const (
	TEST = iota
)

func Make(t LevelType) types.Level {
	switch t {
	case TEST:
		var b = new(baseLevel)
		for x := 0; x < b.XWidth(); x++ {
			for y := 0; y < b.YWidth(); y++ {
				if x == 0 || x == b.XWidth()-1 {
					b[x][y].cellType = types.VWALL
				} else if y == 0 || y == b.YWidth()-1 {
					b[x][y].cellType = types.HWALL
				} else {
					b[x][y].cellType = types.FLOOR
				}
			}
		}
		return b
	default:
		return nil
	}
}
