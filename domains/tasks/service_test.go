package tasks_test

import (
	"testing"
	"time"

	"todo-backend/domains/tasks"
	"todo-backend/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateTaskService(t *testing.T) {
	t.Run("test create service", func(t *testing.T) {
		taskRepoMock := new(mocks.TaskRepoMock)

		task := tasks.Task{
			ID:          "4f19cbbc-8c2c-49dd-b48a-eabafb6ab7f2",
			Title:       "test",
			Description: "test",
			Status:      "IN_PROGRESS",
			Image:       "MTExMQ==",
			CreateAt:    time.Date(2023, time.August, 19, 22, 1, 46, 785911500, time.Local),
		}
		taskRepoMock.On("CreateTask", mock.AnythingOfType("*tasks.Task")).Return(&task, nil)
		service := tasks.NewService(taskRepoMock)
		_, err := service.CreateTask(&task)
		t.Run("test store data with no error", func(t *testing.T) {
			assert.Equal(t, nil, err)
		})

	})

}
func TestUpdateTaskService(t *testing.T) {
	t.Run("test update service", func(t *testing.T) {
		taskRepoMock := new(mocks.TaskRepoMock)
		task := tasks.Task{
			ID:          "4f19cbbc-8c2c-49dd-b48a-eabafb6ab7f2",
			Title:       "test",
			Description: "test",
			Status:      "IN_PROGRESS",
			Image:       "MTExMQ==",
			CreateAt:    time.Date(2023, time.August, 19, 22, 1, 46, 785911500, time.Local),
		}
		taskRepoMock.On("UpdateTask", mock.AnythingOfType("*tasks.Task"), mock.AnythingOfType("string")).Return(&task, nil)
		service := tasks.NewService(taskRepoMock)
		_, err := service.UpdateTask(&task, task.ID)
		t.Run("test store data with no error", func(t *testing.T) {
			assert.Equal(t, nil, err)
		})
	})
}

func TestGetTaskByIDService(t *testing.T) {
	t.Run("test get by id service", func(t *testing.T) {
		taskRepoMock := new(mocks.TaskRepoMock)
		task := tasks.Task{
			ID:          "4f19cbbc-8c2c-49dd-b48a-eabafb6ab7f2",
			Title:       "test",
			Description: "test",
			Status:      "IN_PROGRESS",
			Image:       "MTExMQ==",
			CreateAt:    time.Date(2023, time.August, 19, 22, 1, 46, 785911500, time.Local),
		}
		taskRepoMock.On("GetTaskByID", mock.AnythingOfType("string")).Return(&task, nil)
		service := tasks.NewService(taskRepoMock)
		_, err := service.GetTaskByID("4f19cbbc-8c2c-49dd-b48a-eabafb6ab7f2")
		t.Run("test store data with no error", func(t *testing.T) {
			assert.Equal(t, nil, err)
		})
	})
}

func TestDeleteTaskService(t *testing.T) {
	t.Run("test delete service", func(t *testing.T) {
		taskRepoMock := new(mocks.TaskRepoMock)

		taskRepoMock.On("DeleteTask", mock.AnythingOfType("string")).Return(nil)
		service := tasks.NewService(taskRepoMock)
		err := service.DeleteTask("4f19cbbc-8c2c-49dd-b48a-eabafb6ab7f2")
		t.Run("test store data with no error", func(t *testing.T) {
			assert.Equal(t, nil, err)
		})
	})
}
func TestGetListTaskService(t *testing.T) {
	t.Run("test get list service", func(t *testing.T) {
		var results []tasks.Task
		taskRepoMock := new(mocks.TaskRepoMock)
		taskRepoMock.On("GetListTask", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(&results, nil)
		service := tasks.NewService(taskRepoMock)
		_, err := service.GetListTask("", "", "", "")
		t.Run("test store data with no error", func(t *testing.T) {
			assert.Equal(t, nil, err)
		})
	})
}
