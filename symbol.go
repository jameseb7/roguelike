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
