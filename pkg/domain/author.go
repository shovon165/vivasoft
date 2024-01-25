package domain

import (
	"book-crud/pkg/models"
	"book-crud/pkg/types"
)

type IAuthorRepo interface {
	GetAuthors(request map[string]string) ([]models.AuthorDetail, error)
	GetAuthor(authorID uint) (*models.AuthorDetail, error)
	CreateAuthor(author *models.AuthorDetail) error
	UpdateAuthor(author *models.AuthorDetail) error
	DeleteAuthor(authorID uint) error
}

type IAuthorService interface {
	GetAuthors(request map[string]string) ([]types.ReadAuthorResponse, error)
	GetAuthor(authorID uint) (*types.ReadAuthorResponse, error)
	CreateAuthor(request *types.CreateAuthorRequest) error
	UpdateAuthor(authorID uint, request *types.UpdateAuthorRequest) error
	DeleteAuthor(authorID uint) error
}
