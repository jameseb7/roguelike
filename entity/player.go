package entity

type Player {
	id EntityID
	currentAction *Action
}

func (p Player) ID() {
	return p.id
}

func (p *Player) NextAction() (action Action) { 
	if currentAction == nil {
		return PlayerAction{}
	}
	action := *p.currentAction
	p.currentAction = nil
}

func (p *Player) SetAction(action Action) {
	*p.currentAction = action
}

func NewPlayer() Entity {
	p := new(Player)
	p.id = NewEntityID()
}