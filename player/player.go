package player

import "github.com/jameseb7/roguelike/types"

type Player struct {
	CurrentLevel types.Level
	x, y         int
	currentAction types.Action
	actionChannel <-chan types.Action
}

func (p *Player) X() int                       { return p.x }
func (p *Player) Y() int                       { return p.y }
func (p *Player) SetX(x int)                   { p.x = x }
func (p *Player) SetY(y int)                   { p.y = y }
func (p *Player) Symbol() types.Symbol         { return types.PLAYER }
func (p *Player) Parent() types.Level { return p.CurrentLevel}
func (p *Player) SetParent(parent types.Level) { p.CurrentLevel = parent }
func (p *Player) SetAction(a types.Action) {
	p.currentAction = a
}

func SetQuitCallback(callback func()) {
	quit = callback
}

func SetStopCallback(callback func()) {
	stop = callback
}

var quit func()
var stop func()

//Player implements types.Actor
func (p *Player) Act() {
	if p.currentAction == nil {
		stop()
		return 
	}	
	
	complete := p.currentAction.DoAction()
	if complete {
		p.currentAction = nil //let the action be garbage collected
	}
	
}