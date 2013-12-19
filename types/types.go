package types

type Symbol uint64

const ( //Terrain symbols
	BLANK Symbol = iota
	FLOOR
	HWALL
	VWALL
)
const ( //Monster symbols
	PLAYER Symbol = (1 << 16) + iota
)

type Direction uint8

const (
	NORTH Direction = iota
	SOUTH
	EAST
	WEST

	NORTHEAST
	NORTHWEST
	SOUTHEAST
	SOUTHWEST
)

type Level interface {
	SymbolAt(x, y int) Symbol

	XWidth() int
	YWidth() int

	IsOccupied(x, y int) bool

	Put(e Entity, x, y int)
}

type Entity interface {
	X() int
	Y() int

	SetX(x int)
	SetY(y int)

	Symbol() Symbol

	SetParent(parent Level)
}
