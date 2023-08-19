package tasks

import (
	"net/http"
	domainError "todo-backend/domains/errors"
	"todo-backend/domains/tasks"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	Service tasks.TaskService
}

// GetListTask  godoc
//
// @Summary  get list task
// @Description Getting list task from database
// @Accept   json
// @Produce  json
// @Success  200   {object} ListResponse
// @Router   /tasks [get]
func (handler *TaskHandler) GetListTask(c *gin.Context) {
	sortTitle := c.Query("SortTitle")
	sortCreateAt := c.Query("SortCreateAt")
	sortStatus := c.Query("SortStatus")
	searchTitle := c.Query("Title")
	searchDescription := c.Query("Description")
	results, err := handler.Service.GetListTask(sortTitle, sortCreateAt, sortStatus, searchTitle, searchDescription)
	if err != nil {
		c.Error(err)
		return
	}
	var responseItems = make([]TaskResponse, len(results))

	for i, element := range results {
		responseItems[i] = *toResponseModel(&element)
	}

	response := &ListResponse{
		Data: responseItems,
	}

	c.JSON(http.StatusOK, response)
}

// CreateTask  godoc
//
// @Summary  create task
// @Description Adding new task to database
// @Accept   json
// @Produce  json
// @Param   task body  TaskValidator true "Task Data"
// @Success  201   {object} ListResponse
// @Router   /tasks [post]
func (handler *TaskHandler) CreateTask(c *gin.Context) {
	task, err := Bind(c)
	if err != nil {
		appError := domainError.NewAppError(err, domainError.ValidationError)
		c.Error(appError)
		return
	}
	newTask, err := handler.Service.CreateTask(task)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusCreated, *toResponseModel(newTask))
}

// GetTaskByID  godoc
//
// @Summary get task by id
// @Description getting task by id from database
// @Accept json
// @Product json
// @Param id path string true "id"
// @Success  200   {object} ListResponse
// @Router /tasks/{id} [get]
func (handler *TaskHandler) GetTaskByID(c *gin.Context) {
	id := c.Param("id")
	result, err := handler.Service.GetTaskByID(id)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, toResponseModel(result))

}

// UpdateTask  godoc
//
// @Summary update task
// @Description update task to database
// @Accept json
// @Product json
// @Param   task body  TaskValidator true "Task Data"
// @Param id path string true "id"
// @Success  200   {object} ListResponse
// @Router /tasks/{id} [put]
func (handler *TaskHandler) UpdateTask(c *gin.Context) {
	id := c.Param("id")
	task, err := Bind(c)
	if err != nil {
		appError := domainError.NewAppError(err, domainError.ValidationError)
		c.Error(appError)
		return
	}
	task.ID = id
	result, err := handler.Service.GetTaskByID(id)
	if err != nil {
		c.Error(err)
		return
	}
	result.Title = task.Title
	result.Description = task.Description
	result.Image = task.Image
	result.Status = task.Status
	updatedTask, err := handler.Service.UpdateTask(result, id)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, toResponseModel(updatedTask))

}

// DeleteTask  godoc
//
// @Summary delete task
// @Description delete task from database
// @Accept json
// @Product json
// @Param id path string true "id"
// @Success  200   {object} ListResponse
// @Router /tasks/{id} [delete]
func (handler *TaskHandler) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	err := handler.Service.DeleteTask(id)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, "Success")
}
func NewRoutesFactory(group *gin.RouterGroup) func(service tasks.TaskService) {

	taskRoutesFactory := func(service tasks.TaskService) {
		handler := TaskHandler{Service: service}
		group.GET("/", handler.GetListTask)
		group.POST("/", handler.CreateTask)
		group.GET("/:id", handler.GetTaskByID)
		group.PUT("/:id", handler.UpdateTask)
		group.DELETE("/:id", handler.DeleteTask)
	}
	return taskRoutesFactory
}
