package tasks

type TaskRepository interface {
	CreateTask(*Task) (*Task, error)
	GetTaskByID(string) (*Task, error)
	GetListTask(string, string, string, string) ([]Task, error)
	UpdateTask(*Task, string) (*Task, error)
	DeleteTask(string) error
}
