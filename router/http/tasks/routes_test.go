package tasks_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
	domains "todo-backend/domains/tasks"
	"todo-backend/mocks"
	"todo-backend/router/http/tasks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var mockTask = tasks.TaskValidator{
	Title:       "test",
	Description: "test",
	Status:      "IN_PROGRESS",
	Image:       "MTExMQ==",
}

func TestCreateApi(t *testing.T) {
	t.Run("test create api success", func(t *testing.T) {
		task := domains.Task{
			ID:          "4f19cbbc-8c2c-49dd-b48a-eabafb6ab7f2",
			Title:       "test",
			Description: "test",
			Status:      "IN_PROGRESS",
			Image:       "MTExMQ==",
			CreateAt:    time.Date(2023, time.August, 19, 22, 1, 46, 785911500, time.Local),
		}
		taskServiceMock := new(mocks.TaskServiceMock)
		taskServiceMock.On("CreateTask", mock.AnythingOfType("*tasks.Task")).Return(&task, nil)
		gin := gin.New()
		rec := httptest.NewRecorder()
		handler := tasks.TaskHandler{
			Service: taskServiceMock,
		}
		gin.POST("/tasks", handler.CreateTask)

		body, err := json.Marshal(mockTask)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/tasks", strings.NewReader(string(body)))
		gin.ServeHTTP(rec, req)

		t.Run("test status code and response body", func(t *testing.T) {
			assert.Equal(t, http.StatusCreated, rec.Code)

		})

	})

}
func TestUpdateApi(t *testing.T) {
	t.Run("test update api success", func(t *testing.T) {
		task := domains.Task{
			ID:          "4f19cbbc-8c2c-49dd-b48a-eabafb6ab7f2",
			Title:       "test",
			Description: "test",
			Status:      "IN_PROGRESS",
			Image:       "MTExMQ==",
			CreateAt:    time.Date(2023, time.August, 19, 22, 1, 46, 785911500, time.Local),
		}
		taskServiceMock := new(mocks.TaskServiceMock)
		taskServiceMock.On("GetTaskByID", mock.AnythingOfType("string")).Return(&task, nil)
		taskServiceMock.On("UpdateTask", mock.AnythingOfType("*tasks.Task"), mock.AnythingOfType("string")).Return(&task, nil)
		gin := gin.New()
		rec := httptest.NewRecorder()
		handler := tasks.TaskHandler{
			Service: taskServiceMock,
		}
		gin.PUT("/tasks/:id", handler.UpdateTask)

		body, err := json.Marshal(mockTask)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPut, `/tasks/`+task.ID, strings.NewReader(string(body)))
		gin.ServeHTTP(rec, req)

		t.Run("test status code and response body", func(t *testing.T) {
			assert.Equal(t, http.StatusOK, rec.Code)

		})
	})
}
