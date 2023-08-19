package tasks_test

import (
	"testing"
	"time"

	domains "todo-backend/domains/tasks"
	"todo-backend/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateTask(t *testing.T) {
	t.Run("test create service", func(t *testing.T) {
		taskRepoMock := new(mocks.TaskRepoMock)

		task := domains.Task{
			ID:          "4f19cbbc-8c2c-49dd-b48a-eabafb6ab7f2",
			Title:       "test",
			Description: "test",
			Status:      "IN_PROGRESS",
			Image:       "MTExMQ==",
			CreateAt:    time.Date(2023, time.August, 19, 22, 1, 46, 785911500, time.Local),
		}
		taskRepoMock.On("CreateTask", mock.AnythingOfType("*tasks.Task")).Return(&task, nil)
		service := domains.NewService(taskRepoMock)
		results, _ := service.CreateTask(&task)
		t.Run("test store data with no error", func(t *testing.T) {
			assert.Equal(t, "test", results.Title)
		})

	})

}
