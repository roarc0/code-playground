package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/matryer/is"
	"github.com/rs/zerolog/log"

	"github.com/roarc0/go-task-service/internal/api/controllers"
	"github.com/roarc0/go-task-service/internal/config"
	"github.com/roarc0/go-task-service/internal/models"
)

func TestApiCreate(t *testing.T) {
	is := is.New(t)
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	t.Cleanup(cancel)

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	}))

	cfg := &defaultConfig

	go func() {
		taskController := controllers.NewDefaultTaskController(ctx, cfg)
		srv, err := NewAPI(cfg, taskController.Handler(), &log.Logger)
		is.NoErr(err)
		if err := srv.Run(ctx); err != nil {
			is.NoErr(err)
		}
	}()

	time.Sleep(50 * time.Millisecond)

	taskCreate := models.TaskCreate{
		Method: "GET",
		URL:    mockServer.URL,
		Headers: http.Header{
			"Content-Type": []string{"application/text"},
		},
	}

	taskCreated := testCreateTask(is, cfg, taskCreate)

	taskResult := testGetTaskUntilDone(t, cfg, taskCreated.ID, 1*time.Second)

	expectedTaskResult := models.TaskResult{
		ID:             taskCreated.ID,
		Status:         models.StatusDone,
		HTTPStatusCode: 200,
		Headers:        nil,
		Length:         2,
	}
	is.Equal(taskResult.Headers.Get("Content-Type"), "text/plain; charset=utf-8")
	is.Equal(taskResult.Headers.Get("Content-Length"), "2")
	taskResult.Headers = nil
	is.True(reflect.DeepEqual(expectedTaskResult, *taskResult))
}

func testGetTaskUntilDone(t *testing.T, cfg *config.Config, id string, timeoutDuration time.Duration) *models.TaskResult {
	is := is.New(t)

	var timeout <-chan time.Time
	var taskResult *models.TaskResult
	var err error

	running := true
	for running {
		select {
		case <-timeout:
			t.Fatalf("timeout reached while waiting for response")
			running = false
		default:
			taskResult, err = testGetTask(is, cfg, id)
			is.NoErr(err)
			is.True(taskResult != nil)

			if taskResult.Status == models.StatusDone {
				running = false
			} else {
				time.Sleep(30 * time.Millisecond)
			}
		}
		timeout = time.After(timeoutDuration)
	}
	return taskResult
}

func testGetTask(is *is.I, cfg *config.Config, id string) (*models.TaskResult, error) {
	var buf bytes.Buffer
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://%s/task/%s", cfg.Addr(), id), &buf)
	is.NoErr(err)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var taskResult models.TaskResult
	err = json.NewDecoder(res.Body).Decode(&taskResult)
	is.NoErr(err)

	return &taskResult, nil
}

func testCreateTask(is *is.I, cfg *config.Config, taskCreate models.TaskCreate) models.TaskCreated {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(taskCreate)
	is.NoErr(err)

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://%s/task", cfg.Addr()), &buf)
	is.NoErr(err)

	res, err := http.DefaultClient.Do(req)
	is.NoErr(err)

	var taskCreated models.TaskCreated
	err = json.NewDecoder(res.Body).Decode(&taskCreated)
	is.NoErr(err)

	_, err = uuid.Parse(taskCreated.ID)
	is.NoErr(err)

	return taskCreated
}

var defaultConfig = config.Config{
	ListenAddress: "localhost",
	Port:          8080,
	TaskStore: config.StoreConfig{
		Type: "memory",
	},
}
