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
func (p *Player) SetActionChannel(ac <-chan types.Action) { p.actionChannel = ac }

//Player implements types.Actor
func (p *Player) Act() {
	if p.currentAction == nil {
		if p.actionChannel == nil{
			panic("nil actionChannel for Player")
		}

		p.currentAction = <-p.actionChannel
		//if ok == false {
		//	panic("closed actionChannel for Player")
		//}
		p.currentAction.SetPatient(p)
	}	
	
	complete := p.currentAction.DoAction()
	if complete == true {
		p.currentAction = nil //let the action be garbage collected
	}
	
}