package entity

type ItemType int
const (
	None ItemType = iota << 16 //flags indicating special item statuses can be stored in the lower 16 bits
	Stone
)

type Stackable interface {
	Entity
	Add(n int)
	Count() int
}

type stack struct {
	count int
}
func (s *stack) Add(n int) {s.count += n}
func (s *stack) Count() int {return s.count}
func newStack() *stack {
	s := new(stack)
	s.count = 1
	return s
}

func NewItem(t ItemType) Entity{
	switch t {
	case Stone:
		e := new(stone)
		e.id = NewEntityID()
		e.stack = newStack()
		return e
	default:
		return nil
	}
}