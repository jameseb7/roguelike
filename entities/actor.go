package entites

type Action interface {
	ActionType() //no-op to distinguish Actions
}

type PlayerAction struct{}
func (PlayerAction) ActionType(){}

type MoveAction struct {
	dir Direction
}
func (MoveAction) ActionType(){}


type Actor interface {
	Entity
	NextAction() Action
}
