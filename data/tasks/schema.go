package tasks

import "time"

type Task struct {
	ID          string `gorm:"primary_key"`
	Title       string
	Description string
	Status      string
	Image       string
	CreateAt    time.Time
}
