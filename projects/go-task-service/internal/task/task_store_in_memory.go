package task

import (
	"context"
	"errors"
	"sync"

	"github.com/roarc0/go-task-service/internal/models"
)

// StoreInMemory is a very simple example implementation for a task store.
type StoreInMemory struct {
	sync.RWMutex
	tasks      map[string]*models.Task
	taskRunner Runner
	updateChan chan models.TaskResult
}

// NewTaskStoreInMemory creates a non production ready task store used for testing or development.
func NewTaskStoreInMemory(taskRunner Runner) (*StoreInMemory, error) {
	return &StoreInMemory{
		tasks:      make(map[string]*models.Task),
		taskRunner: taskRunner,
		updateChan: make(chan models.TaskResult),
	}, nil
}

// Start the Task Store
func (ts *StoreInMemory) Start(ctx context.Context) {
	go taskResultListener(ctx, ts.updateChan, ts.UpdateTask)
}

// Close implements io.Closer
func (ts *StoreInMemory) Close() error {
	return nil
}

// GetTask returns the current task status
func (ts *StoreInMemory) GetTask(taskID string) (*models.TaskResult, error) {
	ts.RLock()
	defer ts.RUnlock()

	task, exists := ts.tasks[taskID]
	if !exists {
		return nil, ErrTaskNotFound
	}

	taskResultCopy := task.TaskResult // no need for deep copies here.
	return &taskResultCopy, nil
}

// CreateTask creates a new task given the TaskCreate
func (ts *StoreInMemory) CreateTask(taskCreate models.TaskCreate) (*models.TaskCreated, error) {
	newTask := models.NewTask(taskCreate)

	ts.Lock()
	defer ts.Unlock()

	_, exists := ts.tasks[newTask.TaskResult.ID]
	if exists {
		// This shouldn't happen since the id is an uuid but we're still careful.
		return nil, errors.New("task already exists")
	}

	ts.tasks[newTask.TaskResult.ID] = &newTask

	err := ts.taskRunner.Run(newTask, ts.updateChan)
	if err != nil {
		return nil, err
	}

	return &models.TaskCreated{ID: newTask.TaskResult.ID}, nil
}

// UpdateTask is used to update the task with a new task status
func (ts *StoreInMemory) UpdateTask(taskID string, updatedTaskResult models.TaskResult) error {
	ts.Lock()
	defer ts.Unlock()

	prevTask, exists := ts.tasks[taskID]
	if !exists {
		return ErrTaskNotFound
	}

	if prevTask.TaskResult.Status == models.StatusError || prevTask.TaskResult.Status == models.StatusDone {
		return ErrTaskAlreadyFinished
	}

	ts.tasks[taskID].TaskResult = updatedTaskResult
	return nil
}

// DeleteTask is used to remove the task from the map.
func (ts *StoreInMemory) DeleteTask(taskID string) error {
	ts.Lock()
	defer ts.Unlock()

	_, exists := ts.tasks[taskID]
	if !exists {
		return ErrTaskNotFound
	}
	delete(ts.tasks, taskID)
	return nil
}
