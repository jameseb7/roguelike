package level

import "container/list"

import "github.com/jameseb7/roguelike/entity"
import "github.com/jameseb7/roguelike/symbol"

type LevelType int

const (
	Empty LevelType = iota
	Test
)

func New(lt LevelType) Level {
	newLevel := new(baseLevel)

	newLevel.entities = make(map[entity.ID] *entityMetadata, 100)
	newLevel.actors = list.New()

	switch lt {
	case Empty:
		//all cells blank is the zero value for level
	case Test:
		for x := 1; x < XWidth-1; x++ {
			for y := 1; y < YWidth-1; y++ {
				newLevel.cells[x][y].baseSymbol = symbol.Floor
			}
		}
		for y := 0; y < YWidth; y++ {
			newLevel.cells[0][y].baseSymbol = symbol.VWall
			newLevel.cells[XWidth-1][y].baseSymbol = symbol.VWall
		}
		for x := 0; x < XWidth; x++ {
			newLevel.cells[x][0].baseSymbol = symbol.HWall
			newLevel.cells[x][YWidth-1].baseSymbol = symbol.HWall
		}
	}
	
	return newLevel
}