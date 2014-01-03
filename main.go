package main

import "math/rand"
import "time"

import "github.com/jameseb7/roguelike/regions"
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
var quit = false

func main() {
	rand.Seed(time.Now().Unix())
	
	initCurses()
	defer endCurses()

	var r = regions.Make(regions.TEST)
	p = new(player.Player)
	r.Level(0).Put(p, 40, 10)
	r.Level(0).AddActor(p)

	player.SetStopCallback(setStop)

	drawCurrentLevel()

	for !quit {
		switch ch := C.getch(); ch {
		case C.KEY_UP, 'k', '8':
			runCommand(1, MOVE, types.NORTH)
		case C.KEY_DOWN, 'j', '2':
			runCommand(1, MOVE, types.SOUTH)
		case C.KEY_RIGHT, 'l', '6':
			runCommand(1, MOVE, types.EAST)
		case C.KEY_LEFT, 'h', '4':
			runCommand(1, MOVE, types.WEST)
		case 'y', '7':
			runCommand(1, MOVE, types.NORTHWEST)
		case 'u', '9':
			runCommand(1, MOVE, types.NORTHEAST)
		case 'm', '3':
			runCommand(1, MOVE, types.SOUTHEAST)
		case 'n', '1':
			runCommand(1, MOVE, types.SOUTHWEST)
		case '<':
			runCommand(1, MOVE, types.UP)
		case '>':
			runCommand(1, MOVE, types.DOWN)
		case 'q':
			quit = true
		}
		drawCurrentLevel()
	}

}
