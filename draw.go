package main

import "github.com/jameseb7/roguelike/symbol"
import "github.com/jameseb7/roguelike/level"

//#cgo LDFLAGS: -lncurses
//#include <curses.h>
import "C"

func drawSymbol(x, y int, s symbol.Symbol) {
	cx, cy := C.int(x), C.int(y)
	switch s {
	case symbol.Blank:
		C.mvaddch(cy, cx, ' ')
	case symbol.Floor:
		C.mvaddch(cy, cx, '.')
	case symbol.HWall:
		C.mvaddch(cy, cx, '-')
	case symbol.VWall:
		C.mvaddch(cy, cx, '|')
	case symbol.Player:
		C.mvaddch(cy, cx, '@')
	default:
		C.mvaddch(cy, cx, '\000')
	}
}

func drawLevel(l level.Level) {
	for x := 0; x < level.XWidth; x++ {
		for y := 0; y < level.YWidth; y++ {
			drawSymbol(x, y, l.SymbolAt(x, y))
		}
	}
}

func drawCurrentLevel() {
	drawLevel(currentLevel)
}
