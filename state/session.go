package state

import "time"

type Session struct {
	id       Uuid
	duration time.Duration
	workout  Uuid
	metrics  []*Metric
	notes    string
	created  time.Time
	upated   time.Time
}
