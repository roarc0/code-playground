package task

import (
	"context"
	"errors"
	"io"

	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"

	"github.com/roarc0/go-task-service/models"
)

var (
	//ErrTaskNotFound is used when the task is not found
	ErrTaskNotFound = errors.New("task not found")
	//ErrTaskAlreadyExists is used when the task already exists
	ErrTaskAlreadyExists = errors.New("task already exists")
	//ErrTaskAlreadyFinished is used when the task update failed
	ErrTaskAlreadyFinished = errors.New("task already finished")
)

// Store is used to manage the tasks
type Store interface {
	io.Closer
	Start(context.Context)
	CreateTask(taskCreate models.TaskCreate) (*models.TaskCreated, error)
	GetTask(taskID string) (*models.TaskResult, error)
	UpdateTask(taskID string, updatedTaskResult models.TaskResult) error
	DeleteTask(taskID string) error
}

// StoreConfig is used to configure the
type StoreConfig struct {
	Type   string `yaml:"type"`
	Params any    `yaml:"params,omitempty"`
}

// StoreFactory builds the taskstore given the configuration
func StoreFactory(cfg StoreConfig, runner Runner) (Store, error) {
	switch cfg.Type {
	case "badger":
		b, err := yaml.Marshal(cfg.Params)
		if err != nil {
			return nil, err
		}
		var c StoreBadgerConfig
		err = yaml.Unmarshal(b, &c)
		if err != nil {
			return nil, err
		}
		return NewTaskStoreBadger(c, runner)
	case "memory":
		return NewTaskStoreInMemory(runner)
	default:
		return nil, errors.New("unsupported task store type")
	}
}

func taskResultListener(ctx context.Context, updateChan <-chan models.TaskResult, updateFn func(string, models.TaskResult) error) {
	log.Info().Msg("Task result listener starting...")
	for {
		select {
		case <-ctx.Done():
			log.Info().Msg("Task result listener shutting down...")
			return
		case taskResult := <-updateChan:
			log.Info().Any("taskID", taskResult.ID).Msg("Received task update")
			if err := updateFn(taskResult.ID, taskResult); err != nil {
				log.Error().Err(err).Msg("Task update error")
			}
		}
	}
}
