package tasks

import domains "todo-backend/domains/tasks"

func ToDBModel(entity *domains.Task) *Task {
	return &Task{
		ID:          entity.ID,
		Title:       entity.Title,
		Description: entity.Description,
		Image:       entity.Image,
		Status:      entity.Status,
		CreateAt:    entity.CreateAt,
	}
}
func ToDomainModel(entity *Task) *domains.Task {
	return &domains.Task{
		ID:          entity.ID,
		Title:       entity.Title,
		Description: entity.Description,
		Image:       entity.Image,
		Status:      entity.Status,
		CreateAt:    entity.CreateAt,
	}
}
