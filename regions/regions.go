package regions

import "math/rand"
import "github.com/jameseb7/roguelike/types"


type RegionType int
const (
	TEST = iota
)

func Make(t RegionType) types.Region {
	switch t {
	case TEST:
		r := new(testRegion)
		n := rand.Intn(10) + 1
		r.lvls = make([]types.Level, n)
		return r
	default:
		return nil
	}
}
