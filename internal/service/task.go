package service

import (
	"context"
	"sync"
	"time"

	"testovoe/internal/models"
	"testovoe/internal/storage"

	"github.com/google/uuid"
)

type TaskService struct {
	storage *storage.MemoryTaskStorage
}

func NewTaskService(storage *storage.MemoryTaskStorage) *TaskService {
	return &TaskService{storage: storage}
}

func (s *TaskService) CreateTask(ctx context.Context, data interface{}) (*models.Task, error) {
	task := &models.Task{
		ID:        generateID(),
		Data:      data,
		Status:    models.StatusPending,
		CreatedAt: time.Now(),
	}

	if err := s.storage.Create(task); err != nil {
		return nil, err
	}

	go s.processTask(task)

	return task, nil
}

func (s *TaskService) GetTask(ctx context.Context, id string) (*models.Task, error) {
	return s.storage.Get(id)
}

func (s *TaskService) DeleteTask(ctx context.Context, id string) error {
	return s.storage.Delete(id)
}

func (s *TaskService) processTask(task *models.Task) {
	task.Mu.Lock()
	task.Status = models.StatusProcessing
	now := time.Now()
	task.StartedAt = &now
	s.storage.Update(task)
	task.Mu.Unlock()

	var wg sync.WaitGroup
	steps := 100
	wg.Add(steps)

	for i := 0; i < steps; i++ {
		go func(step int) {
			defer wg.Done()

			time.Sleep(time.Second * 2)

			task.Mu.Lock()
			task.Progress = float64(step+1) / float64(steps) * 100
			s.storage.Update(task)
			task.Mu.Unlock()
		}(i)
	}

	wg.Wait()

	task.Mu.Lock()
	defer task.Mu.Unlock()

	task.Status = models.StatusCompleted
	now = time.Now()
	task.CompletedAt = &now
	task.Result = "task completed successfully"
	s.storage.Update(task)
}

func generateID() string {
	return uuid.New().String()
}
