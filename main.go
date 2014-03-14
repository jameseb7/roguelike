package main

import "math/rand"
import "time"

import (
	"github.com/jameseb7/roguelike/entity"
	"github.com/jameseb7/roguelike/level"
	"github.com/jameseb7/roguelike/action"
	"github.com/jameseb7/roguelike/direction"
)



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

var player *entity.Player
var currentLevel level.Level
var quit = false

func main() {
	rand.Seed(time.Now().Unix())
	
	initCurses()
	defer endCurses()

	currentLevel = level.New(level.Empty)
	player = entity.NewPlayer()
	currentLevel.Put(player, 40, 13)

	drawCurrentLevel()

	for !quit {
		var a action.Action

		switch ch := C.getch(); ch {
		case C.KEY_UP, 'k', '8':
			a = action.Move{direction.North}
		case C.KEY_DOWN, 'j', '2':
			a = action.Move{direction.South}
		case C.KEY_RIGHT, 'l', '6':
			a = action.Move{direction.East}
		case C.KEY_LEFT, 'h', '4':
			a = action.Move{direction.West}
		case 'y', '7':
			a = action.Move{direction.NorthWest}
		case 'u', '9':
			a = action.Move{direction.NorthEast}
		case 'm', '3':
			a = action.Move{direction.SouthEast}
		case 'n', '1':
			a = action.Move{direction.SouthWest}
		case 'q':
			quit = true
		}

		player.SetAction(a)
		currentLevel.Run()
		drawCurrentLevel()
	}

}
