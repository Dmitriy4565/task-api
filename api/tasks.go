package api

import (
	"fmt"
	"net/http"

	"testovoe/internal/models"
	"testovoe/internal/service"
	"testovoe/pkg/prot"

	"github.com/gin-gonic/gin"
)

func createTaskHandler(s *service.TaskService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request models.CreateTaskRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			prot.Error(c, http.StatusBadRequest, "Invalid request body")
			return
		}

		task, err := s.CreateTask(c.Request.Context(), request.Data)
		if err != nil {
			prot.Error(c, http.StatusInternalServerError, err.Error())
			return
		}

		prot.Success(c, http.StatusCreated, models.TaskResponse{
			ID:        task.ID,
			Status:    task.Status,
			CreatedAt: task.CreatedAt,
		})
	}
}

func getTaskStatusHandler(s *service.TaskService) gin.HandlerFunc {
	return func(c *gin.Context) {
		taskID := c.Param("id")

		task, err := s.GetTask(c.Request.Context(), taskID)
		if err != nil {
			prot.Error(c, http.StatusNotFound, "Task not found")
			return
		}

		createdAt := task.CreatedAt.Format("2006-01-02 15:04")

		var durationStr string
		if task.CompletedAt != nil {
			duration := int(task.CompletedAt.Sub(task.CreatedAt).Seconds())
			durationStr = fmt.Sprintf("%d sec", duration)
		}

		prot.Success(c, http.StatusOK, models.TaskStatusResponse{
			ID:        task.ID,
			Status:    task.Status,
			CreatedAt: createdAt,
			Duration:  durationStr,
			Result:    task.Result,
			Progress:  task.Progress,
		})
	}
}

func deleteTaskHandler(s *service.TaskService) gin.HandlerFunc {
	return func(c *gin.Context) {
		taskID := c.Param("id")

		if err := s.DeleteTask(c.Request.Context(), taskID); err != nil {
			prot.Error(c, http.StatusNotFound, "Task not found")
			return
		}

		prot.Success(c, http.StatusNoContent, nil)
	}
}
