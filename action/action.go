package action

type Action interface {
	ActionType() //no-op to distinguish Actions
}

type Player struct{}
func (Player) ActionType(){}

type Move struct {
	Dir Direction
}
func (Move) ActionType(){}
