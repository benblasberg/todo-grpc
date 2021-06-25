package data

import (
	"errors"
	"sync"

	"github.com/pborman/uuid"
)

//MemStore is an in-memory implementation of the todo datastore interface
type MemStore struct {
	data map[string]*Task
	mu   sync.Mutex
}

func NewInMemoryDataStore() *MemStore {
	return &MemStore{data: make(map[string]*Task)}
}

func (m *MemStore) AddTodo(task *Task) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, ok := m.data[task.Id]; ok {
		return errors.New("Task with id " + task.Id + " already exists")
	}

	task.Id = uuid.New()

	m.data[task.Id] = task
	return nil
}

func (m *MemStore) GetTasks() (*[]Task, error) {
	var tasks []Task

	m.mu.Lock()
	defer m.mu.Unlock()

	for _, t := range m.data {
		tasks = append(tasks, Task{Id: t.Id, Title: t.Title, Description: t.Description})
	}

	return &tasks, nil
}
