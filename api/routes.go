package api

import (
	"testovoe/internal/service"

	"github.com/gin-gonic/gin"
)

func SetupTaskRoutes(router *gin.Engine, taskService *service.TaskService) {
	api := router.Group("/api/v1")
	{
		tasks := api.Group("/tasks")
		{
			tasks.POST("", createTaskHandler(taskService))

			tasks.GET("/:id", getTaskStatusHandler(taskService))

			tasks.DELETE("/:id", deleteTaskHandler(taskService))
		}
	}
}
