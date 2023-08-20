package mocks

import (
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
	/*result := tasks.Task{
		ID:          "4f19cbbc-8c2c-49dd-b48a-eabafb6ab7f2",
		Title:       "test",
		Description: "test",
		Status:      "IN_PROGRESS",
		Image:       "MTExMQ==",
		CreateAt:    time.Date(2023, time.August, 19, 22, 1, 46, 785911500, time.Local),
	}*/
	argument := m.Called(task)
	var r0 *tasks.Task
	if argument.Get(0) != nil {
		r0 = argument.Get(0).(*tasks.Task)
	}
	var r1 error
	if argument.Get(1) != nil {
		r1 = argument.Get(1).(error)
	}
	return r0, r1
}

func (m *TaskServiceMock) UpdateTask(task *tasks.Task, id string) (*tasks.Task, error) {
	/*result := tasks.Task{
		ID:          "4f19cbbc-8c2c-49dd-b48a-eabafb6ab7f2",
		Title:       "test",
		Description: "test",
		Status:      "IN_PROGRESS",
		Image:       "MTExMQ==",
		CreateAt:    time.Date(2023, time.August, 19, 22, 1, 46, 785911500, time.Local),
	}*/
	argument := m.Called(task, id)
	var r0 *tasks.Task
	if argument.Get(0) != nil {
		r0 = argument.Get(0).(*tasks.Task)
	}
	var r1 error
	if argument.Get(1) != nil {
		r1 = argument.Get(1).(error)
	}
	return r0, r1
}

func (m *TaskServiceMock) GetTaskByID(id string) (*tasks.Task, error) {
	/*result := tasks.Task{
		ID:          "4f19cbbc-8c2c-49dd-b48a-eabafb6ab7f2",
		Title:       "test",
		Description: "test",
		Status:      "IN_PROGRESS",
		Image:       "MTExMQ==",
		CreateAt:    time.Date(2023, time.August, 19, 22, 1, 46, 785911500, time.Local),
	}*/
	argument := m.Called(id)
	var r0 *tasks.Task
	if argument.Get(0) != nil {
		r0 = argument.Get(0).(*tasks.Task)
	}
	var r1 error
	if argument.Get(1) != nil {
		r1 = argument.Get(1).(error)
	}
	return r0, r1
}
func (m *TaskServiceMock) DeleteTask(id string) error {

	err := m.Called(id).Error(0)
	if err != nil {
		return err
	}
	return nil
}
