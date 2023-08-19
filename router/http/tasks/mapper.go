package tasks

import domains "todo-backend/domains/tasks"

func toResponseModel(entity *domains.Task) *TaskResponse {
	return &TaskResponse{
		ID:          entity.ID,
		Title:       entity.Title,
		Description: entity.Description,
		Image:       entity.Image,
		Status:      entity.Status,
		CreateAt:    entity.CreateAt,
	}
}
