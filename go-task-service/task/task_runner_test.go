package task

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/matryer/is"
	"github.com/roarc0/go-task-service/models"
)

func TestRunNewTask(t *testing.T) {
	is := is.New(t)

	mockResponse := "ok"
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(mockResponse))
		w.WriteHeader(http.StatusOK)
	}))
	defer mockServer.Close()
	tr := &taskRunnerImpl{httpClient: mockServer.Client()}

	updateChan := make(chan models.TaskResult, 1)
	id := uuid.NewString()
	task := models.Task{
		TaskCreate: models.TaskCreate{Method: "GET", URL: mockServer.URL},
		TaskResult: models.TaskResult{ID: id, Status: models.StatusNew},
	}
	err := tr.Run(task, updateChan)
	is.NoErr(err)

	// Check if the TaskResult is sent through the update channel
	waitForTaskUpdate(t, &task, updateChan)
	is.Equal(id, task.TaskResult.ID)
	is.Equal(models.StatusInProcess, task.TaskResult.Status)

	// Check the next update
	waitForTaskUpdate(t, &task, updateChan)
	task.TaskResult.Headers = nil
	is.Equal(task.TaskResult, models.TaskResult{
		ID:             id,
		Status:         models.StatusDone,
		HTTPStatusCode: 200,
		Headers:        nil,
		Length:         int64(len(mockResponse)),
	})

}

func TestRunNewTaskError(t *testing.T) {
	is := is.New(t)

	tr := &taskRunnerImpl{httpClient: http.DefaultClient}

	updateChan := make(chan models.TaskResult, 1)
	id := uuid.NewString()
	task := models.Task{
		TaskCreate: models.TaskCreate{Method: "GET", URL: "http://a.b.c.nonexistent.xy1a"},
		TaskResult: models.TaskResult{ID: id, Status: models.StatusNew},
	}
	err := tr.Run(task, updateChan)
	is.NoErr(err)

	// Check if the TaskResult is sent through the update channel
	waitForTaskUpdate(t, &task, updateChan)
	is.Equal(id, task.TaskResult.ID)
	is.Equal(models.StatusInProcess, task.TaskResult.Status)

	// Check the next update
	waitForTaskUpdate(t, &task, updateChan)
	is.Equal(task.TaskResult, models.TaskResult{
		ID:             id,
		Status:         models.StatusError,
		HTTPStatusCode: 0,
		Headers:        nil,
		Length:         0,
	})

}

func waitForTaskUpdate(t *testing.T, task *models.Task, updateChan <-chan models.TaskResult) {
	select {
	case taskResult := <-updateChan:
		task.TaskResult = taskResult
	case <-time.After(1 * time.Second):
		t.Error("Timed out waiting for TaskResult to be sent through the channel")
	}
}
