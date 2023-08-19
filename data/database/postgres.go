package db

import (
	"fmt"
	"todo-backend/config"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

// Connect to a database handle from a connection string.
func Connect(configuration *config.Database) (*gorm.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", configuration.Host, configuration.Port, configuration.DB, configuration.User, configuration.Password)
	db, err := gorm.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	return db, nil
}
