package actions

import "github.com/jameseb7/roguelike/types"

type Move struct {
	Dir types.Direction
	patient types.Entity
}

func (m *Move) SetPatient(e types.Entity) {
	m.patient = e
}

func (m *Move) DoAction() (complete bool) {
	m.patient.Parent().Move(m.patient, m.Dir)
	return true
}