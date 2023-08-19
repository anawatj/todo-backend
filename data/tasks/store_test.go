package tasks_test

import (
	"database/sql"
	"testing"
	"time"
	"todo-backend/data/tasks"
	domains "todo-backend/domains/tasks"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/jinzhu/gorm"
)

type v1Suite struct {
	db   *gorm.DB
	mock sqlmock.Sqlmock
	task domains.Task
}

func TestCreateTask(t *testing.T) {
	s := &v1Suite{}

	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()

	if err != nil {
		t.Errorf("Failed to open mock sql db, got error: %v", err)
	}

	if db == nil {
		t.Error("mock db is null")
	}

	if s.mock == nil {
		t.Error("sqlmock is null")
	}

	s.db, err = gorm.Open("postgres", db)
	if err != nil {
		t.Errorf("Failed to open gorm db, got error: %v", err)
	}

	if s.db == nil {
		t.Error("gorm db is null")
	}
	id := uuid.New()
	task := domains.Task{
		ID:          id.String(),
		Title:       "test",
		Description: "test",
		Image:       "MTExMQ==",
		CreateAt:    time.Now(),
		Status:      "IN_PROGRESS",
	}
	store := &tasks.Store{
		DB: s.db,
	}
	s.mock.ExpectBegin()
	store.CreateTask(&task)
	s.mock.ExpectCommit()
	assert.Equal(t, "test", task.Title)
	defer db.Close()

}
func TestGetTaskByID(t *testing.T) {
	s := &v1Suite{}

	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()

	if err != nil {
		t.Errorf("Failed to open mock sql db, got error: %v", err)
	}

	if db == nil {
		t.Error("mock db is null")
	}

	if s.mock == nil {
		t.Error("sqlmock is null")
	}

	s.db, err = gorm.Open("postgres", db)
	if err != nil {
		t.Errorf("Failed to open gorm db, got error: %v", err)
	}

	if s.db == nil {
		t.Error("gorm db is null")
	}
	id := uuid.New()
	task := domains.Task{
		ID:          id.String(),
		Title:       "test",
		Description: "test",
		Image:       "MTExMQ==",
		CreateAt:    time.Now(),
		Status:      "IN_PROGRESS",
	}

	store := &tasks.Store{
		DB: s.db,
	}

	s.mock.ExpectBegin()
	store.CreateTask(&task)
	rows := sqlmock.NewRows([]string{"ID", "Title", "Description", "Image", "CreateAt", "Status"}).AddRow(task.ID, task.Title, task.Description, task.CreateAt, task.Image, task.Status)
	s.mock.ExpectQuery(`SELECT`).WillReturnRows(rows)
	store.GetTaskByID(id.String())
	s.mock.ExpectCommit()
	defer db.Close()
}

func TestGetListTask(t *testing.T) {
	s := &v1Suite{}

	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()

	if err != nil {
		t.Errorf("Failed to open mock sql db, got error: %v", err)
	}

	if db == nil {
		t.Error("mock db is null")
	}

	if s.mock == nil {
		t.Error("sqlmock is null")
	}

	s.db, err = gorm.Open("postgres", db)
	if err != nil {
		t.Errorf("Failed to open gorm db, got error: %v", err)
	}

	if s.db == nil {
		t.Error("gorm db is null")
	}
	id := uuid.New()
	task := domains.Task{
		ID:          id.String(),
		Title:       "test",
		Description: "test",
		Image:       "MTExMQ==",
		CreateAt:    time.Now(),
		Status:      "IN_PROGRESS",
	}

	store := &tasks.Store{
		DB: s.db,
	}

	s.mock.ExpectBegin()
	store.CreateTask(&task)
	rows := sqlmock.NewRows([]string{"ID", "Title", "Description", "Image", "CreateAt", "Status"}).AddRow(task.ID, task.Title, task.Description, task.CreateAt, task.Image, task.Status)
	s.mock.ExpectQuery(`SELECT`).WillReturnRows(rows)
	store.GetListTask("", "", "", "", "")
	s.mock.ExpectCommit()
	defer db.Close()
}

func TestUpdateTask(t *testing.T) {
	s := &v1Suite{}

	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()

	if err != nil {
		t.Errorf("Failed to open mock sql db, got error: %v", err)
	}

	if db == nil {
		t.Error("mock db is null")
	}

	if s.mock == nil {
		t.Error("sqlmock is null")
	}

	s.db, err = gorm.Open("postgres", db)
	if err != nil {
		t.Errorf("Failed to open gorm db, got error: %v", err)
	}

	if s.db == nil {
		t.Error("gorm db is null")
	}
	id := uuid.New()
	task := domains.Task{
		ID:          id.String(),
		Title:       "test",
		Description: "test",
		Image:       "MTExMQ==",
		CreateAt:    time.Now(),
		Status:      "IN_PROGRESS",
	}

	store := &tasks.Store{
		DB: s.db,
	}

	s.mock.ExpectBegin()
	store.CreateTask(&task)
	rows := sqlmock.NewRows([]string{"ID", "Title", "Description", "Image", "CreateAt", "Status"}).AddRow(task.ID, task.Title, task.Description, task.CreateAt, task.Image, task.Status)
	s.mock.ExpectQuery(`SELECT`).WillReturnRows(rows)
	store.GetTaskByID(id.String())
	task.Title = "Test11"
	store.UpdateTask(&task, task.ID)
	s.mock.ExpectCommit()
	defer db.Close()
}

func TestDeleteTask(t *testing.T) {
	s := &v1Suite{}

	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()

	if err != nil {
		t.Errorf("Failed to open mock sql db, got error: %v", err)
	}

	if db == nil {
		t.Error("mock db is null")
	}

	if s.mock == nil {
		t.Error("sqlmock is null")
	}

	s.db, err = gorm.Open("postgres", db)
	if err != nil {
		t.Errorf("Failed to open gorm db, got error: %v", err)
	}

	if s.db == nil {
		t.Error("gorm db is null")
	}
	id := uuid.New()
	task := domains.Task{
		ID:          id.String(),
		Title:       "test",
		Description: "test",
		Image:       "MTExMQ==",
		CreateAt:    time.Now(),
		Status:      "IN_PROGRESS",
	}

	store := &tasks.Store{
		DB: s.db,
	}

	s.mock.ExpectBegin()
	store.CreateTask(&task)
	rows := sqlmock.NewRows([]string{"ID", "Title", "Description", "Image", "CreateAt", "Status"}).AddRow(task.ID, task.Title, task.Description, task.CreateAt, task.Image, task.Status)
	s.mock.ExpectQuery(`SELECT`).WillReturnRows(rows)
	store.GetTaskByID(id.String())
	store.DeleteTask(task.ID)
	s.mock.ExpectCommit()
	defer db.Close()
}
