package entities

type ID uint64

type Entity interface {
	EntityID() ID
}

var currentID uint64 = 0

func NewEntityID() ID {
	currentID++
	if currentID == 0 {
		panic("Out of Entity IDs")
	}
	return currentID
}
