package registry

import "context"

type TaskFunc func(ctx context.Context, payload string) error

var taskHandlers = make(map[string]TaskFunc)

func RegisterTask(name string, handler TaskFunc) {
	taskHandlers[name] = handler
}

func GetTaskHandler(name string) TaskFunc {
	return taskHandlers[name]
}
