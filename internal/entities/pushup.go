package entities

import "time"

type PushupOperationType string

// Values based on table pushups.type CHECK constraint
const (
	OpAdd      PushupOperationType = "add"
	OpSubtract PushupOperationType = "subtract"
)

type Pushup struct {
	Id          int
	Repetitions int
	Type        PushupOperationType
	CreatedAt   time.Time
}

type NewPushup struct {
	Repetitions int
	Type        PushupOperationType
	CreatedAt   time.Time
}
