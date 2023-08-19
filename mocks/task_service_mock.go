package mocks

import (
	"time"
	"todo-backend/domains/tasks"

	"github.com/stretchr/testify/mock"
)

type TaskServiceMock struct {
	mock.Mock
	taskService tasks.TaskService
}

func (m *TaskServiceMock) GetListTask(sort string, sortType string, searchTitle string, searchDescription string) ([]tasks.Task, error) {
	var results []tasks.Task
	argument := m.Called(sort, sortType, searchTitle, searchDescription)
	if argument.Is("error") {
		return nil, argument.Error(0)
	}
	return results, nil
}
func (m *TaskServiceMock) CreateTask(task *tasks.Task) (*tasks.Task, error) {
	result := tasks.Task{
		ID:          "4f19cbbc-8c2c-49dd-b48a-eabafb6ab7f2",
		Title:       "test",
		Description: "test",
		Status:      "IN_PROGRESS",
		Image:       "MTExMQ==",
		CreateAt:    time.Date(2023, time.August, 19, 22, 1, 46, 785911500, time.Local),
	}
	argument := m.Called(task)
	if argument.Is("error") {
		return nil, argument.Error(0)
	}
	return &result, nil
}

func (m *TaskServiceMock) UpdateTask(task *tasks.Task, id string) (*tasks.Task, error) {
	result := tasks.Task{
		ID:          "4f19cbbc-8c2c-49dd-b48a-eabafb6ab7f2",
		Title:       "test",
		Description: "test",
		Status:      "IN_PROGRESS",
		Image:       "MTExMQ==",
		CreateAt:    time.Date(2023, time.August, 19, 22, 1, 46, 785911500, time.Local),
	}
	argument := m.Called(task, id)
	if argument.Is("error") {
		return nil, argument.Error(0)
	}
	return &result, nil
}

func (m *TaskServiceMock) GetTaskByID(id string) (*tasks.Task, error) {
	result := tasks.Task{
		ID:          "4f19cbbc-8c2c-49dd-b48a-eabafb6ab7f2",
		Title:       "test",
		Description: "test",
		Status:      "IN_PROGRESS",
		Image:       "MTExMQ==",
		CreateAt:    time.Date(2023, time.August, 19, 22, 1, 46, 785911500, time.Local),
	}
	argument := m.Called(id)
	if argument.Is("error") {
		return nil, argument.Error(0)
	}
	return &result, nil
}
func (m *TaskServiceMock) DeleteTask(id string) error {

	err := m.Called(id).Error(0)
	if err != nil {
		return err
	}
	return nil
}
