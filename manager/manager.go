package manager

import (
	"fmt"

	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
	"github.com/warlck/go-orchestrator/task"
)

type Manager struct {
	Pending       queue.Queue
	TaskDb        map[string][]task.Task
	EventDb       map[string][]task.TaskEvent
	Workers       []string
	WorkerTaskMap map[string][]uuid.UUID
	TaskWokerMap  map[uuid.UUID]string
}

func (m *Manager) SelectWorker() {
	fmt.Println("SelectWorker placehoder")
}

func (m *Manager) UpdateTasks() {
	fmt.Println("UpdateTasks placeholder")
}

func (m *Manager) SendWork() {
	fmt.Println("SendWork placeholder")
}
