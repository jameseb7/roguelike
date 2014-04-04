package main

import "fmt"
import "unsafe"

import "github.com/jameseb7/roguelike/symbol"
import "github.com/jameseb7/roguelike/level"

//#cgo LDFLAGS: -lncurses
//#include <curses.h>
//#include <stdlib.h>
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
	case symbol.Rock:
		C.mvaddch(cy, cx, '#')
	case symbol.Player:
		C.mvaddch(cy, cx, '@')
	case symbol.Stone:
		C.mvaddch(cy, cx, '*')
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

func drawInfoBar() {
	cx, cy := C.int(0), C.int(level.YWidth)
	C.mvaddstr(cy, cx, C.CString("T: "))
	C.addstr(C.CString(fmt.Sprint(currentLevel.Turn())))
}

func displayInventory(){
	var str *C.char
	C.clear()
	C.move(0,0)
	str = C.CString("Inventory contents:\n")
	C.addstr(str)
	C.free(unsafe.Pointer(str))
	for i, v := range inventory {
		if inventory[i] != nil {
			C.addch(C.chtype(inventoryChar(byte(i))))
			str = C.CString(" - ")
			C.addstr(str)
			C.free(unsafe.Pointer(str))
			str = C.CString(v.name)
			C.addstr(str)
			C.free(unsafe.Pointer(str))
			C.addch(C.chtype('\n'))
		}
	}
}