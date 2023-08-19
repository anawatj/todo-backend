package tasks

import "time"

const (
	IN_PROGRESS = "IN_PROGRESS"
	COMPLETE    = "COMPLETE"
)

type Task struct {
	ID          string
	Title       string
	Description string
	Status      string
	Image       string
	CreateAt    time.Time
}
