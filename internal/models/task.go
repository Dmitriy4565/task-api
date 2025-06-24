package models

import (
	"sync"
	"time"
)

type TaskStatus string

const (
	StatusPending    TaskStatus = "pending"
	StatusProcessing TaskStatus = "processing"
	StatusCompleted  TaskStatus = "completed"
	StatusFailed     TaskStatus = "failed"
)

type Task struct {
	ID          string
	Data        interface{}
	Status      TaskStatus
	Result      interface{}
	Error       error
	Progress    float64
	CreatedAt   time.Time
	StartedAt   *time.Time
	CompletedAt *time.Time
	Mu          sync.Mutex
}

type CreateTaskRequest struct {
	Data interface{} `json:"data"`
}

type TaskResponse struct {
	ID        string     `json:"id"`
	Status    TaskStatus `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
}

type TaskStatusResponse struct {
	ID        string      `json:"id"`
	Status    TaskStatus  `json:"status"`
	CreatedAt string      `json:"created_at"`
	Duration  string      `json:"duration,omitempty"`
	Result    interface{} `json:"result,omitempty"`
	Progress  float64     `json:"progress"`
}
