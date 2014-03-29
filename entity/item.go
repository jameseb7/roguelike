package entity

type ItemType int
const (
	None ItemType = iota
	Stone
)

func NewItem(t ItemType) Entity{
	switch t {
	case Stone:
		e := new(stone)
		e.id = NewEntityID()
		return e
	default:
		return nil
	}
}