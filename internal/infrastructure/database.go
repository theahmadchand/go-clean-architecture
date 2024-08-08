package infrastructure

import (
	"github.com/theahmadchand/go-clean-architecture/internal/entities"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDatabaseConnection() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&entities.Post{})
	if err != nil {
		return nil, err
	}

	return db, nil
}