package entity

import "github.com/jameseb7/roguelike/symbol"

type stone struct{
	id ID
}

func (s stone) EntityID() ID {return s.id}
func (s stone) EntityName() string {return "A stone"}
func (s stone) EntitySymbol() symbol.Symbol {return symbol.Stone}
func (s stone) Size() int {return Small}