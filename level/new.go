package level

import "list"

type LevelType
const (
	Empty LevelType = iota
)

func New(lt int) Level {
	newLevel := new(baseLevel)

	switch levelType {
	case Empty:
		//all cells blank is the zero value for level
		newLevel.entities = make(map[entity.ID] *entityMetadata, 100)
		newLevel.actors = list.New()
	}
}