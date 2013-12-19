package player

import "github.com/jameseb7/roguelike/types"

type Player struct {
	currentLevel types.Level
	x, y         int
}

func (p *Player) X() int                       { return p.x }
func (p *Player) Y() int                       { return p.y }
func (p *Player) SetX(x int)                   { p.x = x }
func (p *Player) SetY(y int)                   { p.y = y }
func (p *Player) Symbol() types.Symbol         { return types.PLAYER }
func (p *Player) SetParent(parent types.Level) { p.currentLevel = parent }
