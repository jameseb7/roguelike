package main

import "github.com/jameseb7/roguelike/levels"

//#cgo LDFLAGS: -lncurses
//#include <curses.h>
import "C"

func initCurses() {
	C.initscr()
	C.cbreak()
	C.noecho()
	C.nonl()
	C.intrflush(C.stdscr, true)
	C.keypad(C.stdscr, false)
}

func endCurses() {
	C.nocbreak()
	C.echo()
	C.nl()
	C.endwin()
}

func main() {
	initCurses()
	defer endCurses()

	var l = levels.Make(levels.TEST)
	for x := 0; x < l.XWidth(); x++ {
		for y := 0; y < l.YWidth(); y++ {
			drawSymbol(x, y, l.SymbolAt(x, y))
		}
	}
	C.getch()
}
