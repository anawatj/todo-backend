package tasks

import "time"

type TaskResponse struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	Image       string    `json:"image"`
	CreateAt    time.Time `json:"createAt"`
}
type ListResponse struct {
	Data []TaskResponse `json:"data"`
}
