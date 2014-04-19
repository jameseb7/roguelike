package main

import "fmt"
import "unsafe"

import "github.com/jameseb7/roguelike/symbol"
import "github.com/jameseb7/roguelike/level"
import "github.com/jameseb7/roguelike/entity"

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

func displayPickUpChoice() []entity.ID {
	var str *C.char
	C.clear()
	C.move(0,0)
	str = C.CString("Pick up what?\n")
	C.addstr(str)
	C.free(unsafe.Pointer(str))

	px, py := currentLevel.EntityLocation(player.EntityID())
	itemsAvailable := currentLevel.ItemsAt(px, py)
	for i, eid := range itemsAvailable {
		C.addch(C.chtype(inventoryChar(byte(i))))
		str = C.CString(" - ")
		C.addstr(str)
		C.free(unsafe.Pointer(str))
		str = C.CString(currentLevel.EntityByID(eid).EntityName())
		C.addstr(str)
		C.free(unsafe.Pointer(str))
	}

	itemsChosen := make([]bool, len(itemsAvailable))
	for {
		ch := C.getch()
		if ch == C.KEY_ENTER || ch == ' ' || ch == '\n'  {
			break
		}
		if ch > C.int(255) {
			continue
		}
		if i := inventoryIndex(byte(ch)); (int(i) < len(itemsChosen)) && 
			(i != 255) {
			if itemsChosen[i] {
				itemsChosen[i] = false
				C.mvaddch(C.int(i+1), 2, C.chtype('-'))
			} else {
				itemsChosen[i] = true
				C.mvaddch(C.int(i+1), 2, C.chtype('+'))
			}
		}
	}
	
	result := make([]entity.ID, 0, len(itemsAvailable))
	for i, v := range itemsChosen {
		if v {
			result = append(result, itemsAvailable[i])
		}
	}
	return result
}