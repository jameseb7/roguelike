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

type Level interface {
	SymbolAt(x, y int) Symbol

	XWidth() int
	YWidth() int

	IsOccupied(x, y int) bool

	Put(e Entity, x, y int) (ok bool)
	Move(e Entity, dir Direction) (ok bool)
}

type Entity interface {
	X() int
	Y() int

	SetX(x int)
	SetY(y int)

	Symbol() Symbol

	SetParent(parent Level)
}
