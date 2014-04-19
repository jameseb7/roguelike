package entity

import "fmt"
import "github.com/jameseb7/roguelike/symbol"

type stone struct{
	*stack
	id ID
}

func (s stone) EntityID() ID {return s.id}
func (s stone) EntityName() string {
	if s.Count() < 2 {
		return "A stone"
	} else {
		return fmt.Sprint(s.Count(), " stones")
	}
}
func (s stone) EntitySymbol() symbol.Symbol {return symbol.Stone}
func (s stone) EntityItemType() ItemType {return Stone}
func (s stone) Size() int {return Small}