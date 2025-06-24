package storage

import (
	"errors"
	"sync"

	"testovoe/internal/models"
)

type MemoryTaskStorage struct {
	mu    sync.RWMutex
	tasks map[string]*models.Task
}

func NewMemoryTaskStorage() *MemoryTaskStorage {
	return &MemoryTaskStorage{
		tasks: make(map[string]*models.Task),
	}
}

func (s *MemoryTaskStorage) Create(task *models.Task) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.tasks[task.ID]; exists {
		return errors.New("task already exists")
	}

	s.tasks[task.ID] = task
	return nil
}

func (s *MemoryTaskStorage) Get(id string) (*models.Task, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	task, exists := s.tasks[id]
	if !exists {
		return nil, ErrTaskNotFound
	}

	return task, nil
}

// ff
func (s *MemoryTaskStorage) Delete(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.tasks[id]; !exists {
		return ErrTaskNotFound
	}

	delete(s.tasks, id)
	return nil
}

func (s *MemoryTaskStorage) Update(task *models.Task) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.tasks[task.ID]; !exists {
		return ErrTaskNotFound
	}

	s.tasks[task.ID] = task
	return nil
}

var ErrTaskNotFound = errors.New("task not found")
