package worker

import (
	"fmt"

	"github.com/warlck/go-orchestrator/task"

	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

type Worker struct {
	Name      string
	Queue     queue.Queue
	Db        map[uuid.UUID]task.Task
	TaskCount int
}

func (w *Worker) CollectStats() {
	fmt.Println("CollectStats placeholder")
}

func (w *Worker) StartTask() {
	fmt.Println("StartTask placeholder")
}

func (w *Worker) StopTask() {
	fmt.Println("Stoptask placeholder")
}

func (w *Worker) RunTask() {
	fmt.Println("RunTask placeholder")
}
