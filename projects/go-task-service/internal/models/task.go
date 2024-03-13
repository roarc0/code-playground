package models

import (
	"net/http"

	"github.com/google/uuid"
)

// TaskCreate is used to create a new task
type TaskCreate struct {
	Method  string      `json:"method"`
	URL     string      `json:"url"`
	Headers http.Header `json:"headers"`
}

// TaskCreated is used to as a response for the TaskCreate
type TaskCreated struct {
	ID string `json:"id"`
}

// Status is an enum that reports the current status of the task
type Status string

const (
	// StatusNew is the initial status for a task
	StatusNew Status = "new"
	// StatusInProcess is used when the task is being processed
	StatusInProcess Status = "in_process"
	//StatusDone is used when the task has finished to execute successfully
	StatusDone Status = "done"
	// StatusError is used when the task has finished to execute with an error
	StatusError Status = "error"
)

// TaskResult contains the information that report the status of the execution
type TaskResult struct {
	ID             string      `json:"id"`
	Status         Status      `json:"status"`
	HTTPStatusCode int         `json:"httpStatusCode"`
	Headers        http.Header `json:"headers"`
	Length         int64       `json:"length"`
}

// Task contains all the information to represent a task.
type Task struct {
	TaskCreate TaskCreate
	TaskResult TaskResult
}

// NewTask returns a new task ready to be executed.
func NewTask(taskCreate TaskCreate) Task {
	newTask := Task{
		TaskCreate: taskCreate,
		TaskResult: TaskResult{
			ID:             uuid.NewString(),
			Status:         StatusNew,
			HTTPStatusCode: 0,
			Headers:        nil,
			Length:         0,
		},
	}
	return newTask
}
