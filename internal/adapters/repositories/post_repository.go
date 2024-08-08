package repositories

import (
	"errors"
	"github.com/theahmadchand/go-clean-architecture/internal/entities"
	"gorm.io/gorm"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (repository *PostRepository) Create(post *entities.Post) error {
	return repository.db.Create(post).Error
}

func (repository *PostRepository) GetAll() ([]entities.Post, error) {
	var posts []entities.Post
	if err := repository.db.Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (repository *PostRepository) GetByID(id string) (*entities.Post, error) {
	var post entities.Post
	if err := repository.db.First(&post, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (repository *PostRepository) Delete(id string) error {
	result := repository.db.Where("id = ?", id).Delete(&entities.Post{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("record not found")
	}
	return nil
}