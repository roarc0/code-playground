package task

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/matryer/is"

	"github.com/roarc0/go-task-service/internal/config"
	"github.com/roarc0/go-task-service/internal/models"
)

func TestTaskStoreFactory(t *testing.T) {
	type args struct {
		cfg config.StoreConfig
	}
	tests := []struct {
		name   string
		args   args
		wantFn func(*testing.T, Store, error)
	}{
		{
			name: "memory",
			args: args{cfg: config.StoreConfig{Type: "memory"}},
			wantFn: func(t *testing.T, ts Store, err error) {
				is := is.New(t)
				is.NoErr(err)

				tsInMemory, ok := ts.(*StoreInMemory)
				is.True(ok)
				is.True(tsInMemory != nil)
			},
		}, {
			name: "badger",
			args: args{cfg: config.StoreConfig{Type: "badger", Params: map[string]any{
				"in-memory": true,
			}}},
			wantFn: func(t *testing.T, ts Store, err error) {
				is := is.New(t)
				is.NoErr(err)

				tsBadger, ok := ts.(*StoreBadger)
				is.True(ok)
				is.True(tsBadger != nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts, err := StoreFactory(tt.args.cfg, NewRunner())
			tt.wantFn(t, ts, err)
		})
	}
}

func TestTaskStore(t *testing.T) {
	type args struct {
		cfg config.StoreConfig
	}
	tests := []struct {
		name string
		fn   func(t *testing.T, ts Store)
		args args
	}{
		{
			name: "memory",
			fn:   testTaskStore200,
			args: args{cfg: config.StoreConfig{Type: "memory"}},
		},
		{
			name: "badger",
			fn:   testTaskStore200,
			args: args{cfg: config.StoreConfig{Type: "badger", Params: map[string]any{
				"in-memory": true,
			}}},
		},
		{
			name: "memory_404",
			fn:   testTaskStore404,
			args: args{cfg: config.StoreConfig{Type: "memory"}},
		},
		{
			name: "badger_404",
			fn:   testTaskStore404,
			args: args{cfg: config.StoreConfig{Type: "badger", Params: map[string]any{
				"in-memory": true,
			}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is := is.New(t)
			ts, err := StoreFactory(tt.args.cfg, NewRunner())
			is.NoErr(err)
			defer ts.Close()
			ts.Start(context.Background())
			tt.fn(t, ts)
		})
	}
}

func testTaskStore200(t *testing.T, ts Store) {
	responseBody := "ok"
	testTaskStore(t, ts, 200, &responseBody)
}

func testTaskStore404(t *testing.T, ts Store) {
	testTaskStore(t, ts, 404, nil)
}

func testTaskStore(t *testing.T, ts Store, httpStatusCode int, responseBody *string) {
	is := is.New(t)
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if responseBody != nil {
			w.Write([]byte(*responseBody))
		}
		w.WriteHeader(httpStatusCode)
	}))
	defer mockServer.Close()

	taskCreated := testTaskStoreCreate(t, ts, mockServer.URL)
	t.Log(taskCreated)

	var taskResult models.TaskResult
	var timeout <-chan time.Time
	timeoutDuration := 1 * time.Second
	running := true
	for running {
		select {
		case <-timeout:
			t.Fatalf("timeout reached while waiting for response")
			running = false
		default:
			taskResult = testTaskStoreGet(t, ts, taskCreated.ID)
			t.Log(taskResult)
			if taskResult.Status == models.StatusDone {
				running = false
			} else {
				time.Sleep(20 * time.Millisecond)
			}
		}
		timeout = time.After(timeoutDuration)
	}

	is.Equal(taskResult.HTTPStatusCode, httpStatusCode)
	is.Equal(taskResult.Status, models.StatusDone)

	testTaskStoreDelete(t, ts, taskCreated.ID)
}

func testTaskStoreCreate(t *testing.T, ts Store, url string) models.TaskCreated {
	is := is.New(t)
	taskCreate := models.TaskCreate{
		Method:  "GET",
		URL:     url,
		Headers: http.Header{},
	}
	taskCreated, err := ts.CreateTask(taskCreate)
	is.NoErr(err)
	is.True(taskCreated != nil && len(taskCreated.ID) > 0)
	return *taskCreated
}

func testTaskStoreGet(t *testing.T, ts Store, taskID string) models.TaskResult {
	is := is.New(t)
	taskResult, err := ts.GetTask(taskID)
	is.NoErr(err)
	is.True(taskResult != nil && len(taskResult.ID) > 0)
	return *taskResult
}

func testTaskStoreDelete(t *testing.T, ts Store, taskID string) {
	is := is.New(t)
	err := ts.DeleteTask(taskID)
	is.NoErr(err)
	_, err = ts.GetTask(taskID)
	is.Equal(ErrTaskNotFound, err)
}
