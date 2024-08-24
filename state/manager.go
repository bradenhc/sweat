package state

type Manager struct {
	store Store
}

type Store struct {
}

func NewManager() *Manager {
	return &Manager{}
}
