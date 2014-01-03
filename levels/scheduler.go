package levels

import "container/list"

import "github.com/jameseb7/roguelike/types"
import "github.com/jameseb7/roguelike/player"

type scheduler struct{
	stop bool
	turn int
	actors *list.List
	currentActor *list.Element
	waitingActors *list.List
	actorSlots map[types.Actor] *list.Element
}

func newScheduler() *scheduler {
	s := new(scheduler)
	s.turn = 1
	s.actors = list.New()
	s.waitingActors = list.New()
	s.actorSlots = make(map[types.Actor] *list.Element, 20)
	return s
}

func (s *scheduler) AddActor(a types.Actor) {
	s.waitingActors.PushBack(a)
}

func (s *scheduler) addActorProper(a types.Actor) {
	e := s.actors.PushBack(a)
	s.actorSlots[a] = e
}

func (s *scheduler) RemoveActor(a types.Actor) {
	e := s.actorSlots[a]
	s.actors.Remove(e)
	delete(s.actorSlots, a)
}

func (s *scheduler) stopCallback() {
	s.stop = true
}

func (s *scheduler) Run() {
	player.SetStopCallback(s.stopCallback)
	s.stop = false

	//check currentActor is initialised
	if s.currentActor == nil {
		s.currentActor = s.actors.Front()
		//if still nil then try moving the waiting actors
		if s.currentActor == nil {
			s.moveWaitingActors()
			s.currentActor = s.actors.Front()
			//if still nil then there are no actors so nothing need be done
			if s.currentActor == nil {
				return
			}
		}
	}


	for { //turn loop
		for { //Actor loop
			s.currentActor.Value.(types.Actor).Act()
			if s.stop {
				return
			}
			s.currentActor = s.currentActor.Next()
		}
		
		s.moveWaitingActors()
		s.turn++
		s.currentActor = s.actors.Front()
	}
}

func (s *scheduler) moveWaitingActors() {
	for e := s.waitingActors.Front(); e != nil; e = e.Next() {
		s.addActorProper(e.Value.(types.Actor))
		s.waitingActors.Remove(e)
	}
}
