package main

import "github.com/jameseb7/roguelike/levels"
import "github.com/jameseb7/roguelike/player"
import "github.com/jameseb7/roguelike/types"

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

	l.Move(p, types.NORTH)
	drawLevel(l)
	C.getch()

	l.Move(p, types.EAST)
	drawLevel(l)
	C.getch()

	l.Move(p, types.SOUTH)
	drawLevel(l)
	C.getch()

	l.Move(p, types.WEST)
	drawLevel(l)
	C.getch()

	l.Move(p, types.NORTHEAST)
	drawLevel(l)
	C.getch()

}
