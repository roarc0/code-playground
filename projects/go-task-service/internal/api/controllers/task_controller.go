package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog/log"

	"github.com/roarc0/go-task-service/internal/config"
	"github.com/roarc0/go-task-service/internal/models"
	"github.com/roarc0/go-task-service/internal/task"
)

type taskCtxKey string

const (
	taskIDKey    = "taskID"
	taskStoreKey = "taskStore"
)

// TaskController implements the HTTP server that makes HTTP request to a third party
type TaskController struct {
	taskStore task.Store
}

func NewDefaultTaskController(ctx context.Context, cfg *config.Config) *TaskController {
	taskRunner := task.NewRunner()
	taskStore, err := task.StoreFactory(cfg.TaskStore, taskRunner)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create task store")
	}
	taskStore.Start(ctx)
	return NewTaskController(taskStore)
}

// NewTaskController creates a new server
func NewTaskController(taskStore task.Store) *TaskController {
	return &TaskController{
		taskStore: taskStore,
	}
}

// Handler returns the http.Handler for the whole task service
func (s *TaskController) Handler() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	// Better a JWT token.
	// r.Use(middleware.BasicAuth("realm", map[string]string{
	// 	"login": "password",
	// }))
	r.Use(s.getTaskStoreCtx)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("🤖"))
	})

	r.Route("/task", func(r chi.Router) {
		r.Post("/", createTask)
		r.Route("/{taskID}", func(r chi.Router) {
			r.Get("/", getTask)
			//r.Delete("/", cancelTask)
		})
	})
	return r
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var taskCreate models.TaskCreate
	if err := json.NewDecoder(r.Body).Decode(&taskCreate); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	taskStore, ok := r.Context().Value(taskCtxKey(taskStoreKey)).(task.Store)
	if !ok {
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}

	taskCreated, err := taskStore.CreateTask(taskCreate)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(taskCreated)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func getTask(w http.ResponseWriter, r *http.Request) {
	taskStore, ok := r.Context().Value(taskCtxKey(taskStoreKey)).(task.Store)
	if !ok {
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}

	taskID := chi.URLParam(r, taskIDKey)
	log.Debug().Str(taskIDKey, taskID).Msg("Getting task")

	taskResult, err := taskStore.GetTask(taskID)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(taskResult)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (s *TaskController) getTaskStoreCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), taskCtxKey(taskStoreKey), s.taskStore)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
