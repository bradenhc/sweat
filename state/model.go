package state

type State struct {
	workouts  map[Uuid]*Workout
	exercises map[Uuid]*Exercise
	sessions  []*Session
}

func NewState() *State {
	return &State{}
}
