package main

import "github.com/jameseb7/roguelike/levels"
import "github.com/jameseb7/roguelike/player"

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
	var p = new(player.Player)
	l.Put(p, 40, 10)

	drawLevel(l)
	C.getch()

}
