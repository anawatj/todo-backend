package tasks

import (
	domainError "todo-backend/domains/errors"
	domains "todo-backend/domains/tasks"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

const (
	createError = "Error in creating new tasks"
	readError   = "Error in finding tasks in the database"
	listError   = "Error in getting tasks from the database"
	updateError = "Error in updating tasks to the database"
	deleteError = "Error in deleting tasks from the database"
)

type Store struct {
	DB *gorm.DB
}

func New(db *gorm.DB) *Store {
	db.AutoMigrate(&Task{})

	return &Store{
		DB: db,
	}
}
func (s *Store) CreateTask(task *domains.Task) (*domains.Task, error) {
	entity := ToDBModel(task)
	err := s.DB.Create(entity).Error
	if err != nil {
		appErr := domainError.NewAppError(errors.Wrap(err, createError), domainError.RepositoryError)
		return nil, appErr
	}

	return ToDomainModel(entity), nil
}
func (s *Store) GetTaskByID(id string) (*domains.Task, error) {
	result := &Task{}
	query := s.DB.Where("id=?", id).First(result)
	if query.RecordNotFound() {
		appErr := domainError.NewAppErrorWithType(domainError.NotFound)
		return nil, appErr
	}

	if err := query.Error; err != nil {
		appErr := domainError.NewAppError(errors.Wrap(err, readError), domainError.RepositoryError)
		return nil, appErr
	}

	return ToDomainModel(result), nil
}
func (s *Store) GetListTask(sort string, sortType string, searchTitle string, searchDescription string) ([]domains.Task, error) {
	var results []Task
	var where string
	var whereParams []string
	var orderBy string
	where = where + " 1=1 "
	if searchDescription != "" {
		where = where + " AND description=? "
		whereParams = append(whereParams, searchDescription)
	}

	if searchTitle != "" {
		where = where + " AND title=? "
		whereParams = append(whereParams, searchTitle)
	}

	orderBy = sort + " " + sortType

	err := s.DB.Order(orderBy).Where(where, whereParams).Find(&results).Error
	if err != nil {
		appErr := domainError.NewAppError(errors.Wrap(err, listError), domainError.RepositoryError)
		return nil, appErr
	}

	var tasks = make([]domains.Task, len(results))

	for i, element := range results {
		tasks[i] = *ToDomainModel(&element)
	}

	return tasks, nil
}
func (s *Store) UpdateTask(task *domains.Task, id string) (*domains.Task, error) {
	var result Task
	query := s.DB.Where("id=?", id).First(result)
	if query.RecordNotFound() {
		appErr := domainError.NewAppErrorWithType(domainError.NotFound)
		return nil, appErr
	}

	if err := query.Error; err != nil {
		appErr := domainError.NewAppError(errors.Wrap(err, updateError), domainError.RepositoryError)
		return nil, appErr
	}
	entity := ToDBModel(task)
	err := s.DB.Update(entity).Error
	if err != nil {
		appErr := domainError.NewAppError(errors.Wrap(err, createError), domainError.RepositoryError)
		return nil, appErr
	}

	return ToDomainModel(entity), nil
}

func (s *Store) DeleteTask(id string) error {

	err := s.DB.Where("id=?", id).Delete(&Task{}).Error
	if err != nil {
		appErr := domainError.NewAppError(errors.Wrap(err, deleteError), domainError.RepositoryError)
		return appErr
	}

	return nil
}
