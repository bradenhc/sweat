package state_test

import (
	"testing"

	"hitchcock.codes/sweat/state"
)

func createTestExercise(t *testing.T) (e *state.Exercise) {
	e, err := state.NewExercise("foo")
	if err != nil {
		t.Errorf("Failed to create exercise: %v", err)
	}
	return
}

func TestNewExercise(t *testing.T) {
	e := createTestExercise(t)
	if e.Name() != "foo" {
		t.Errorf("e.Name() == %v, expected '%v'", e.Name(), "foo")
	}
}

func TestSetExerciseName(t *testing.T) {
	e := createTestExercise(t)
	e.SetName("bar")
	if e.Name() != "bar" {
		t.Errorf("e.Name() == %v, expected '%v'", e.Name(), "bar")
	}
	if e.Updated() == e.Created() {
		t.Error("Exercise update time should not be the same as create time")
	}
}

func TestAddMetricToExercise(t *testing.T) {
	e := createTestExercise(t)
}
