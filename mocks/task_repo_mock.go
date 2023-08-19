package mocks

import (
	"time"
	domains "todo-backend/domains/tasks"

	"github.com/stretchr/testify/mock"
)

type TaskRepoMock struct {
	mock.Mock
	taskRepo domains.TaskRepository
}

func (m *TaskRepoMock) GetListTask(sort string, sortType string, searchTitle string, searchDescription string) ([]domains.Task, error) {
	var results []domains.Task
	err := m.Called(sort, sortType, searchTitle, searchDescription).Error(0)
	if err != nil {
		return nil, err
	}
	return results, nil
}
func (m *TaskRepoMock) CreateTask(task *domains.Task) (*domains.Task, error) {

	result := domains.Task{
		ID:          "4f19cbbc-8c2c-49dd-b48a-eabafb6ab7f2",
		Title:       "test",
		Description: "test",
		Status:      "IN_PROGRESS",
		Image:       "MTExMQ==",
		CreateAt:    time.Date(2023, time.August, 19, 22, 1, 46, 785911500, time.Local),
	}
	err := m.Called(task).Error(0)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
func (m *TaskRepoMock) UpdateTask(task *domains.Task, id string) (*domains.Task, error) {
	var result domains.Task
	err := m.Called(task, id).Error(0)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (m *TaskRepoMock) GetTaskByID(id string) (*domains.Task, error) {
	var result domains.Task
	err := m.Called(id).Error(0)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
func (m *TaskRepoMock) DeleteTask(id string) error {

	err := m.Called(id).Error(0)
	if err != nil {
		return err
	}
	return nil
}
