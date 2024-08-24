package state

import (
	"slices"
	"time"
)

// Exercise holds information about a single exercise, such as its name and the types of metrics
// that can be collected on it.
type Exercise struct {
	id       Uuid
	name     string
	metrics  []string
	created  time.Time
	updated  time.Time
	archived bool
}

// NewExercise creates a new exercise with the provided name and returns it. If an error occurs,
// then the returned exercise pointer will be nil and the returned error will explain what went
// wrong.
func NewExercise(name string) (e *Exercise, err error) {
	id, err := NewUuid()
	if err != nil {
		return
	}

	e = &Exercise{
		id:      id,
		name:    name,
		created: time.Now(),
		updated: time.Now(),
	}

	return
}

func (e *Exercise) SetName(name string) {
	e.name = name
	e.update()
}

func (e *Exercise) Name() string {
	return e.name
}

func (e *Exercise) AddMetric(name string) bool {
	if !slices.Contains(e.metrics, name) {
		e.metrics = append(e.metrics, name)
		e.update()
		return true
	}
	return false
}

func (e *Exercise) Created() time.Time {
	return e.created
}

func (e *Exercise) Updated() time.Time {
	return e.updated
}

func (e *Exercise) Archive() {
	if !e.archived {
		e.archived = true
		e.update()
	}
}

func (e *Exercise) Unarchive() {
	if e.archived {
		e.archived = false
		e.update()
	}
}

func (e *Exercise) IsArchived() bool {
	return e.archived
}

func (e *Exercise) update() {
	e.updated = time.Now()
}
