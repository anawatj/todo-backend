package router

import (
	"net/http"

	"todo-backend/domains/tasks"
	//"github.com/eldimious/golang-api-showcase/domain/books"

	taskRoutes "todo-backend/router/http/tasks"
	//booksRoutes "github.com/eldimious/golang-api-showcase/router/http/books"
	errors "todo-backend/router/http/errors"

	//healthRoutes "github.com/eldimious/golang-api-showcase/router/http/health"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// NewHTTPHandler returns the HTTP requests handler
func NewHTTPHandler(taskSvc tasks.TaskService) http.Handler {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	//config.AddAllowHeaders("Authorization")
	router.Use(cors.New(config))

	router.Use(errors.Handler)

	//healthGroup := router.Group("/health")
	//healthRoutes.NewRoutesFactory()(healthGroup)

	api := router.Group("/api")
	//api.Use(authMiddleware)

	taskGroup := api.Group("/tasks")
	taskRoutes.NewRoutesFactory(taskGroup)(taskSvc)

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//booksRoutes.NewRoutesFactory(authorsGroup)(booksSvc)
	return router
}
