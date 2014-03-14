package level

import "container/list"

import "github.com/jameseb7/roguelike/entity"

type LevelType int

const (
	Empty LevelType = iota
)

func New(lt LevelType) Level {
	newLevel := new(baseLevel)

	switch lt {
	case Empty:
		//all cells blank is the zero value for level
		newLevel.entities = make(map[entity.ID] *entityMetadata, 100)
		newLevel.actors = list.New()
	}
	return newLevel
}