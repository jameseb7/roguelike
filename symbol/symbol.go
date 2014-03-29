package symbol

type Symbol uint64

const(
	Blank = iota
	Floor
	HWall
	VWall
	Rock

	Player
	
	Stone
)