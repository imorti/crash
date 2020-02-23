package service

import (
	"errors"
	"math/rand"

	"github.com/imorti/crash/rest-api/entity"
	"github.com/imorti/crash/rest-api/repository"
)

// Post - to be used in post requests
var (
	repo repository.PostRepository
)

// PostService interface
type PostService interface {
	Validate(post *entity.Post) error
	Create(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}
type service struct{}

// NewPostService - for creating new service
func NewPostService(repoPass repository.PostRepository) PostService {
	repo = repoPass
	return &service{}
}

// Validate - validating a post before saving to data store
func (*service) Validate(post *entity.Post) error {

	if post == nil {
		err := errors.New("The post cannot be empty")
		return err
	}

	if post.Title == "" {
		err := errors.New("The post title cannot be empty")
		return err
	}

	return nil

}

// Create - creating the actual post in the data store post validation
func (*service) Create(post *entity.Post) (*entity.Post, error) {

	post.ID = rand.Int63()
	return repo.Save(post)
}

// FindAll - getting all documents in our collection
func (*service) FindAll() ([]entity.Post, error) {

	return repo.FindAll()

}
