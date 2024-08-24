package state

import "time"

type Workout struct {
	id        Uuid
	name      string
	exercises []*Exercise
	created   time.Time
	updated   time.Time
	archived  bool
}

func NewWorkout(name string) (w *Workout, err error) {
	id, err := NewUuid()
	if err != nil {
		return
	}

	w = &Workout{
		id:        id,
		name:      name,
		exercises: make([]*Exercise, 0),
		created:   time.Now(),
		updated:   time.Now(),
	}

	return
}
