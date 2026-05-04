package entities

import "time"

type Pushup struct {
	Id          int
	Repetitions int
	Type        string
	CreatedAt   time.Time
}

type NewPushup struct {
	Repetitions int
	Type        string
	CreatedAt   time.Time
}
