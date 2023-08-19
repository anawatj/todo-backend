package main

import (
	"net/http"
	"todo-backend/config"
	database "todo-backend/data/database"
	taskStore "todo-backend/data/tasks"
	"todo-backend/docs"
	"todo-backend/domains/tasks"
	router "todo-backend/router/http"
)

// @title   Zog_festiv eCommerce API
// @version  1.0
// @description API for ecommerce website

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host   localhost:8080
// @BasePath  /api

// @schemes http
func main() {
	docs.SwaggerInfo.Title = "Zog_festiv"
	docs.SwaggerInfo.Description = "Yo Yo Yo 148 3 to the 3 to the 6 to the 9 "
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http"}
	configuration, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	// establish DB connection
	db, err := database.Connect(configuration.Database)
	if err != nil {
		panic(err)
	}
	taskRepo := taskStore.New(db)
	taskSrv := tasks.NewService(taskRepo)
	httpRouter := router.NewHTTPHandler(taskSrv)
	err = http.ListenAndServe(":"+configuration.Port, httpRouter)
	if err != nil {
		panic(err)
	}

	defer db.Close()
}
