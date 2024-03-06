package task

import (
	"context"
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/roarc0/go-task-service/models"
)

// Runner executes the task calling a 3rd party service. It reports the updates through a channel.
type Runner interface {
	Run(models.Task, chan<- models.TaskResult) error
}

type taskRunnerImpl struct {
	httpClient *http.Client
}

// RunnerOption is used to set optional features in the task runner instance
type RunnerOption func(tr Runner)

// WithHTTPClient sets the HTTP client for the task runner
func WithHTTPClient(client *http.Client) RunnerOption {
	return func(tr Runner) {
		if trImpl, ok := tr.(*taskRunnerImpl); ok {
			trImpl.httpClient = client
		}
	}
}

func setOptions(tr Runner, options ...RunnerOption) {
	for _, o := range options {
		o(tr)
	}
}

// NewRunner is the component that runs the 3rd party request.
func NewRunner(options ...RunnerOption) Runner {
	tr := &taskRunnerImpl{
		httpClient: &http.Client{
			Timeout: 60 * time.Second,
		},
	}

	setOptions(tr, options...)

	return tr
}

func (tr *taskRunnerImpl) Run(task models.Task, updateChan chan<- models.TaskResult) error {
	if task.TaskResult.Status != models.StatusNew {
		return errors.New("should be in new state")
	}

	go func() {
		log.Info().Any("task", task).Msg("Task created")

		task.TaskResult.Status = models.StatusInProcess
		updateTask(updateChan, &task)

		// TODO We might want to check that we're not calling this service or we could cause some issues.
		// We could set up a white/black list of urls that can be called.

		res, err := tr.HTTPCall(task.TaskCreate)
		if err != nil {
			log.Error().Err(err).Msg("Error in HTTP call")
			task.TaskResult.Status = models.StatusError
			updateTask(updateChan, &task)
			return
		}
		task.TaskResult.Status = models.StatusDone
		task.TaskResult.HTTPStatusCode = res.StatusCode
		task.TaskResult.Headers = res.Header.Clone()

		// We could save the body in some object storage with the ID of the task if needed.
		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Error().Err(err).Msg("Error reading body")
			task.TaskResult.Status = models.StatusError
			updateTask(updateChan, &task)
			return
		}
		defer res.Body.Close()
		task.TaskResult.Length = int64(len(body))

		updateTask(updateChan, &task)
	}()

	return nil
}

func updateTask(updateChan chan<- models.TaskResult, task *models.Task) {
	updateChan <- task.TaskResult
	log.Info().Any("task", task).Msg("Task updated")
}

func (tr *taskRunnerImpl) HTTPCall(taskCreate models.TaskCreate) (*http.Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), tr.httpClient.Timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, taskCreate.Method, taskCreate.URL, nil)
	if err != nil {
		return nil, err
	}
	req.Header = taskCreate.Headers
	return tr.httpClient.Do(req)
}
