package entity

import "github.com/jameseb7/roguelike/symbol"

type Player struct{
	id ID
	currentAction Action
	*inventory
}

func (p Player) EntityID() ID {
	return p.id
}

func (p *Player) EntityName() string {
	return "Player"
}

func (p Player) EntitySymbol() symbol.Symbol {
	return symbol.Player
}

func (p Player) EntityItemType() ItemType {
	return None
}

func (p *Player) Size() int {
	return Large
}

func (p *Player) NextAction(c Context) (a Action) { 
	if p.currentAction == nil {
		return PlayerAction{}
	}
	a = p.currentAction
	p.currentAction = nil
	return
}

func (p *Player) SetAction(a Action) {
	p.currentAction = a
}

func NewPlayer() *Player {
	p := new(Player)
	p.id = NewEntityID()
	p.inventory = newInventory()
	return p
}