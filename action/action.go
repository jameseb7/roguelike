package action

import "github.com/jameseb7/roguelike/direction"

type Action interface {
	ActionType() //no-op to distinguish Actions
}

type Player struct{}
func (Player) ActionType(){}

type Move struct {
	Dir direction.Direction
}
func (Move) ActionType(){}

type Skip struct{}
func (Skip) ActionType(){}