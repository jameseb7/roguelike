package main

import "math/rand"
import "time"
import "log"
import "os"

import (
	"github.com/jameseb7/roguelike/entity"
	"github.com/jameseb7/roguelike/level"
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
	logfile, err := os.Create("roguelike.log")
	if err != nil {
		panic(err)
	}
	log.SetOutput(logfile)

	initCurses()
	defer endCurses()

	currentLevel = level.New(level.Test)
	player = entity.NewPlayer()
	currentLevel.Put(player, 40, 13)

	drawCurrentLevel()
	drawInfoBar()

	for !quit {
		var a entity.Action

		switch ch := C.getch(); ch {
		case C.KEY_UP, 'k', '8':
			a = entity.MoveAction{direction.North}
		case C.KEY_DOWN, 'j', '2':
			a = entity.MoveAction{direction.South}
		case C.KEY_RIGHT, 'l', '6':
			a = entity.MoveAction{direction.East}
		case C.KEY_LEFT, 'h', '4':
			a = entity.MoveAction{direction.West}
		case 'y', '7':
			a = entity.MoveAction{direction.NorthWest}
		case 'u', '9':
			a = entity.MoveAction{direction.NorthEast}
		case 'm', '3':
			a = entity.MoveAction{direction.SouthEast}
		case 'n', '1':
			a = entity.MoveAction{direction.SouthWest}
		case 'i':
			updateInventory()
			displayInventory()
			C.getch()
		case ',':
			items := displayPickUpChoice()
			a = entity.PickUpAction{items}
		case 'q':
			quit = true
		}

		player.SetAction(a)
		currentLevel.Run()
		drawCurrentLevel()
		drawInfoBar()
	}

}
