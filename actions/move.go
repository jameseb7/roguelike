package actions

import "github.com/jameseb7/roguelike/types"

type Move struct {
	Dir types.Direction
	Object types.Entity
}

func (m *Move) DoAction() (complete bool) {
	m.Object.Parent().Move(m.Object, m.Dir)
	return true
}