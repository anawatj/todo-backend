package tasks

type TaskService interface {
	CreateTask(*Task) (*Task, error)
	GetTaskByID(string) (*Task, error)
	GetListTask(string, string, string, string) ([]Task, error)
	UpdateTask(*Task, string) (*Task, error)
	DeleteTask(string) error
}
type Service struct {
	Repository TaskRepository
}

func (svc *Service) CreateTask(task *Task) (*Task, error) {
	return svc.Repository.CreateTask(task)

}
func (svc *Service) GetTaskByID(id string) (*Task, error) {
	return svc.Repository.GetTaskByID(id)
}
func (svc *Service) GetListTask(sortTitle string, sortCreateAt string, searchTitle string, searchDescription string) ([]Task, error) {
	return svc.Repository.GetListTask(sortTitle, sortCreateAt, searchTitle, searchDescription)
}
func (svc *Service) UpdateTask(task *Task, id string) (*Task, error) {
	return svc.Repository.UpdateTask(task, id)
}
func (svc *Service) DeleteTask(id string) error {
	return svc.Repository.DeleteTask((id))
}
func NewService(repository TaskRepository) *Service {
	return &Service{Repository: repository}
}
