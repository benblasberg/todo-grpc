package data

type DataStore interface {
	AddTodo(task *Task) error
	GetTasks() (*[]Task, error)
}

type Task struct {
	Id          string
	Title       string
	Description string
}
