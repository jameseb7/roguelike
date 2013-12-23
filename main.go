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

var p *player.Player


func main() {
	initCurses()
	defer endCurses()

	var l = levels.Make(levels.TEST)
	p = new(player.Player)
	l.Put(p, 40, 10)

	var quit = false
	for !quit {
		drawLevel(l)
		
		switch ch := C.getch(); ch {
		case C.KEY_UP, 'k', '8':
			runCommand(1, MOVE, int(types.NORTH))
		case C.KEY_DOWN, 'j', '2':
			runCommand(1, MOVE, int(types.SOUTH))
		case C.KEY_RIGHT, 'l', '6':
			runCommand(1, MOVE, int(types.EAST))
		case C.KEY_LEFT, 'h', '4':
			runCommand(1, MOVE, int(types.WEST))
		case 'y', '7':
			runCommand(1, MOVE, int(types.NORTHWEST))
		case 'u', '9':
			runCommand(1, MOVE, int(types.NORTHEAST))
		case 'm', '3':
			runCommand(1, MOVE, int(types.SOUTHEAST))
		case 'n', '1':
			runCommand(1, MOVE, int(types.SOUTHWEST))
		case 'q':
			quit = true
		}
	}

}
