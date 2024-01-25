package services

import (
	"book-crud/pkg/domain"
	"book-crud/pkg/models"
	"book-crud/pkg/types"
	"errors"
)

type authorService struct {
	authorRepo domain.IAuthorRepo
	bookRepo   domain.IBookRepo
}

func AuthorServiceInstance(authorRepo domain.IAuthorRepo, bookRepo domain.IBookRepo) domain.IAuthorService {
	return &authorService{
		authorRepo: authorRepo,
		bookRepo:   bookRepo,
	}
}

func (service *authorService) GetAuthors(request map[string]string) ([]types.ReadAuthorResponse, error) {
	authorDetails, err := service.authorRepo.GetAuthors(request)
	if err != nil {
		return nil, err
	}
	if len(authorDetails) == 0 {
		return nil, errors.New("no author found with given query")
	}
	var responses []types.ReadAuthorResponse
	for _, val := range authorDetails {
		responses = append(responses, types.ReadAuthorResponse{
			ID:          val.ID,
			AuthorName:  val.AuthorName,
			Address:     val.Address,
			PhoneNumber: val.PhoneNumber,
		})
	}

	return responses, nil
}

func (service *authorService) GetAuthor(authorID uint) (*types.ReadAuthorResponse, error) {
	authorDetail, err := service.authorRepo.GetAuthor(authorID)
	if err != nil {
		return nil, err
	}

	response := &types.ReadAuthorResponse{
		ID:          authorDetail.ID,
		AuthorName:  authorDetail.AuthorName,
		Address:     authorDetail.Address,
		PhoneNumber: authorDetail.PhoneNumber,
	}

	return response, nil
}

func (service *authorService) CreateAuthor(request *types.CreateAuthorRequest) error {

	authorDetail := &models.AuthorDetail{
		AuthorName:  request.AuthorName,
		Address:     request.Address,
		PhoneNumber: request.PhoneNumber,
	}

	if err := service.authorRepo.CreateAuthor(authorDetail); err != nil {
		return err
	}

	return nil
}

func (service *authorService) UpdateAuthor(authorID uint, request *types.UpdateAuthorRequest) error {
	existingAuthor, err := service.authorRepo.GetAuthor(authorID)
	if err != nil {
		return errors.New("no author found with given author ID")
	}

	if request.AuthorName != "" {
		existingAuthor.AuthorName = request.AuthorName
	}
	if request.Address != "" {
		existingAuthor.Address = request.Address
	}
	if request.PhoneNumber != "" {
		existingAuthor.PhoneNumber = request.PhoneNumber
	}
	if err := service.authorRepo.UpdateAuthor(existingAuthor); err != nil {
		return errors.New("author was not updated")
	}

	return nil
}

// DeleteAuthor handles the delete author request.
func (service *authorService) DeleteAuthor(authorID uint) error {
	if _, err := service.authorRepo.GetAuthor(authorID); err != nil {
		return errors.New("no author found with given author ID")
	}
	if err := service.authorRepo.DeleteAuthor(authorID); err != nil {
		return err
	}
	if err := service.bookRepo.DeleteBooksByAuthorID(authorID); err != nil {
		return err
	}

	return nil
}
