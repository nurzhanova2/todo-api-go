package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/nurzhanova/todo-app/handlers"
)

func SetupRoutes(r *gin.Engine) {
    r.GET("/tasks", handlers.GetTasks)
    r.POST("/tasks", handlers.CreateTask)
    r.PUT("/tasks/:id", handlers.UpdateTask)
    r.DELETE("/tasks/:id", handlers.DeleteTask)
	r.GET("/tasks/:id", handlers.GetTaskByID) 
}
