package task

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/dgraph-io/badger/v4"

	"github.com/roarc0/go-task-service/internal/models"
)

// StoreBadger is an implementation for a task store.
type StoreBadger struct {
	taskRunner Runner
	updateChan chan models.TaskResult
	db         *badger.DB
}

// StoreBadgerConfig is used to configure the badger task store
type StoreBadgerConfig struct {
	InMemory bool   `yaml:"in-memory,omitempty"`
	DbPath   string `yaml:"db-path"`
}

// NewTaskStoreBadger is a task storage system utilizing Badger DB's.
func NewTaskStoreBadger(cfg StoreBadgerConfig, taskRunner Runner) (*StoreBadger, error) {
	var bOpts badger.Options
	if cfg.InMemory {
		bOpts = badger.DefaultOptions("").WithInMemory(true)
	} else {
		bOpts = badger.DefaultOptions(cfg.DbPath)
	}

	db, err := badger.Open(bOpts)
	if err != nil {
		return nil, fmt.Errorf("task-store-badger error: %w", err)
	}

	return &StoreBadger{
		db:         db,
		taskRunner: taskRunner,
		updateChan: make(chan models.TaskResult),
	}, nil
}

// Start the Task Store
func (ts *StoreBadger) Start(ctx context.Context) {
	go taskResultListener(ctx, ts.updateChan, ts.UpdateTask)
}

// Close implements io.Closer
func (ts *StoreBadger) Close() error {
	if ts.db != nil {
		return ts.db.Close()
	}
	return nil
}

// GetTask returns the current task status
func (ts *StoreBadger) GetTask(taskID string) (*models.TaskResult, error) {
	var task models.Task

	err := ts.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(taskID))
		if err != nil {
			if err == badger.ErrKeyNotFound {
				return ErrTaskNotFound
			}
			return err
		}
		return item.Value(func(val []byte) error {
			return json.Unmarshal(val, &task)
		})
	})
	if err != nil {
		return nil, err
	}

	return &task.TaskResult, nil
}

// CreateTask creates a new task given the TaskCreate
func (ts *StoreBadger) CreateTask(taskCreate models.TaskCreate) (*models.TaskCreated, error) {
	newTask := models.NewTask(taskCreate)

	err := ts.db.Update(func(txn *badger.Txn) error {
		_, err := txn.Get([]byte(newTask.TaskResult.ID))
		if err != badger.ErrKeyNotFound {
			return ErrTaskAlreadyExists
		}

		bytes, err := json.Marshal(newTask)
		if err != nil {
			return err
		}

		return txn.Set([]byte(newTask.TaskResult.ID), bytes)
	})
	if err != nil {
		return nil, err
	}

	err = ts.taskRunner.Run(newTask, ts.updateChan)
	if err != nil {
		return nil, err
	}

	return &models.TaskCreated{ID: newTask.TaskResult.ID}, nil
}

// UpdateTask is used to update the task with a new task status
func (ts *StoreBadger) UpdateTask(taskID string, updatedTaskResult models.TaskResult) error {
	return ts.db.Update(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(updatedTaskResult.ID))
		if err != nil {
			if errors.Is(err, badger.ErrKeyNotFound) {
				return ErrTaskNotFound
			}
			return fmt.Errorf("update task error: %w", err)
		}

		var task models.Task
		err = item.Value(func(val []byte) error {
			err = json.Unmarshal(val, &task)
			if err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return err
		}

		if task.TaskResult.Status == models.StatusError || task.TaskResult.Status == models.StatusDone {
			return ErrTaskAlreadyFinished
		}

		task.TaskResult = updatedTaskResult

		newBytes, err := json.Marshal(task)
		if err != nil {
			return err
		}

		return txn.Set([]byte(task.TaskResult.ID), newBytes)
	})
}

// DeleteTask is used to remove the task from the map.
func (ts *StoreBadger) DeleteTask(taskID string) error {
	return ts.db.Update(func(txn *badger.Txn) error {
		_, err := txn.Get([]byte(taskID))
		if err != nil {
			return ErrTaskNotFound
		}

		return txn.Delete([]byte(taskID))
	})
}
