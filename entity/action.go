package entity

import "github.com/jameseb7/roguelike/direction"

type Action interface {
	ActionType() //no-op to distinguish Actions
}

type PlayerAction struct{}
func (PlayerAction) ActionType(){}

type MoveAction struct {
	Dir direction.Direction
}
func (MoveAction) ActionType(){}

type SkipAction struct{}
func (SkipAction) ActionType(){}

type PickUpAction struct {
	Items []ID
}
func (PickUpAction) ActionType(){}

type DropAction struct {
	Item ID
}
func (DropAction) ActionType(){}