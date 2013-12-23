package regions

import "github.com/jameseb7/roguelike/levels"
import "github.com/jameseb7/roguelike/types"

type testRegion struct{
	lvls []types.Level
}

func (tr *testRegion) Level(index int) types.Level {
	if tr.lvls[index] == nil {
		tr.lvls[index] = levels.Make(levels.TEST, tr, index)
	}
	return tr.lvls[index]
}

func (tr *testRegion) NextLevel(current int, dir types.Direction) types.Level {
	var newIndex int
	if dir == types.UP {
		newIndex = current-1
		if newIndex < 0 {
			return tr.Level(current)
		}
	} else if dir == types.DOWN {
		newIndex = current+1
		if newIndex >= len(tr.lvls) {
			return tr.Level(current)
		}
	} else {
		newIndex = current
	}

	return tr.Level(newIndex)
}

func (tr testRegion) Length() int {
	return len(tr.lvls)
}
