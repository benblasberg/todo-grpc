package data

import (
	"errors"
	"sync"
)

//memstore is an in-memory implementation of the todo datastore interface
type memstore struct {
	data map[string]*Task
	mu   sync.Mutex
}

func NewInMemoryDataStore() memstore {
	return memstore{data: make(map[string]*Task)}
}

func (m *memstore) AddTodo(task *Task) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, ok := m.data[task.Id]; ok {
		return errors.New("Task with id " + task.Id + " already exists")
	}

	m.data[task.Id] = task
	return nil
}

func (m *memstore) GetTasks() (*[]Task, error) {
	var tasks []Task

	m.mu.Lock()
	defer m.mu.Unlock()

	for _, t := range m.data {
		tasks = append(tasks, Task{Id: t.Id, Title: t.Title, Description: t.Description})
	}

	return &tasks, nil
}
