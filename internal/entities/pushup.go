package entities

import "time"

type Pushup struct {
	Id           int
	Repetitions  int
	LastModified time.Time
	CreatedAt    time.Time
}

type NewPushup struct {
	Repetitions int
	CreatedAt   time.Time
}
