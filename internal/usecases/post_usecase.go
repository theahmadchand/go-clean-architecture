package usecases

import "github.com/theahmadchand/go-clean-architecture/internal/entities"

type PostRepository interface {
	Create(post *entities.Post) error
	GetAll() ([]entities.Post, error)
	GetByID(id string) (*entities.Post, error)
	Delete(id string) error
}

type PostUseCase struct {
	repo PostRepository
}

func NewPostUseCase(repository PostRepository) *PostUseCase {
	return &PostUseCase{repo: repository}
}

func (useCase *PostUseCase) CreatePost(post *entities.Post) error {
	return useCase.repo.Create(post)
}

func (useCase *PostUseCase) GetPosts() ([]entities.Post, error) {
	return useCase.repo.GetAll()
}

func (useCase *PostUseCase) GetPost(id string) (*entities.Post, error) {
	return useCase.repo.GetByID(id)
}

func (useCase *PostUseCase) DeletePost(id string) error {
	return useCase.repo.Delete(id)
}