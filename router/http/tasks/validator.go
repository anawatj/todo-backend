package tasks

import (
	"time"
	domains "todo-backend/domains/tasks"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TaskValidator struct {
	Title       string `binding:"required" json:"title"`
	Description string `binding:"required" json:"description"`
	Image       string `binding:"required" json:"image"`
	Status      string `binding:"required" json:"status"`
}

func Bind(c *gin.Context) (*domains.Task, error) {
	var json TaskValidator
	if err := c.ShouldBindJSON(&json); err != nil {
		return nil, err
	}
	id := uuid.New()
	task := &domains.Task{
		ID:          id.String(),
		Title:       json.Title,
		Description: json.Description,
		Image:       json.Image,
		Status:      json.Status,
		CreateAt:    time.Now(),
	}
	return task, nil
}
