package types

type Symbol uint64

const ( //Terrain symbols
	BLANK Symbol = iota
	FLOOR
	HWALL
	VWALL
	UPSTAIR
	DOWNSTAIR
)
const ( //Monster symbols
	PLAYER Symbol = (1 << 16) + iota
)

type Region interface {
	Level(index int) Level
	NextLevel(current int, dir Direction) Level
	Length() int
}

type Level interface {
	Index() int

	SymbolAt(x, y int) Symbol

	XWidth() int
	YWidth() int

	IsOccupied(x, y int) bool

	Put(e Entity, x, y int) (ok bool)
	Move(e Entity, dir Direction) (ok bool)

	NextLevel(dir Direction) Level
}

type Entity interface {
	X() int
	Y() int

	SetX(x int)
	SetY(y int)

	Symbol() Symbol

	Parent() Level
	SetParent(parent Level)
}

type Actor interface{
	Act()
}

type Action interface{
	SetPatient(e Entity) //sets an Entity for the actor to act upon
	DoAction() (complete bool)
}
