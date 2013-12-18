package types

type Symbol uint64

const ( //Terrain symbols
	BLANK = iota
	FLOOR
	HWALL
	VWALL
)
const ( //Monster symbols
	PLAYER = (1 << 16) + iota
)

type Level interface {
	SymbolAt(x, y int) Symbol

	XWidth() int
	YWidth() int
}
