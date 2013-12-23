package main

import "github.com/jameseb7/roguelike/types"

//#cgo LDFLAGS: -lncurses
//#include <curses.h>
import "C"

func drawSymbol(x, y int, s types.Symbol) {
	cx, cy := C.int(x), C.int(y)
	switch s {
	case types.BLANK:
		C.mvaddch(cy, cx, ' ')
	case types.FLOOR:
		C.mvaddch(cy, cx, '.')
	case types.HWALL:
		C.mvaddch(cy, cx, '-')
	case types.VWALL:
		C.mvaddch(cy, cx, '|')
	case types.PLAYER:
		C.mvaddch(cy, cx, '@')
	default:
		C.mvaddch(cy, cx, '\000')
	}
}

func drawLevel(l types.Level) {
	for x := 0; x < l.XWidth(); x++ {
		for y := 0; y < l.YWidth(); y++ {
			drawSymbol(x, y, l.SymbolAt(x, y))
		}
	}
}

func drawCurrentLevel() {
	drawLevel(p.CurrentLevel)
}
