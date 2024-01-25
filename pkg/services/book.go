package services

import (
	"book-crud/pkg/domain"
	"book-crud/pkg/models"
	"book-crud/pkg/types"
	"errors"
)

// parent struct to implement interface binding
type bookService struct {
	bookRepo   domain.IBookRepo
	authorRepo domain.IAuthorRepo
}

// interface binding
func BookServiceInstance(bookRepo domain.IBookRepo, authorRepo domain.IAuthorRepo) domain.IBookService {
	return &bookService{
		bookRepo:   bookRepo,
		authorRepo: authorRepo,
	}
}

// all methods of interface are implemented
func (service *bookService) GetBooks(bookID uint) ([]types.BookRequest, error) {
	var allBooks []types.BookRequest
	book := service.bookRepo.GetBooks(bookID)
	if len(book) == 0 {
		return nil, errors.New("No book found")
	}
	for _, val := range book {
		allBooks = append(allBooks, types.BookRequest{
			ID:          val.ID,
			BookName:    val.BookName,
			Author:      val.Author,
			Publication: val.Publication,
		})
	}
	return allBooks, nil
}

func (service *bookService) CreateBook(book *models.BookDetail) error {
	if err := service.bookRepo.CreateBook(book); err != nil {
		return errors.New("BookDetail was not created")
	}
	return nil
}

func (service *bookService) UpdateBook(book *models.BookDetail) error {
	if err := service.bookRepo.UpdateBook(book); err != nil {
		return errors.New("BookDetail update was unsuccessful")
	}
	return nil
}
func (service *bookService) DeleteBook(bookID uint) error {
	if err := service.bookRepo.DeleteBook(bookID); err != nil {
		return errors.New("BookDetail deletion was unsuccessful")
	}
	return nil
}
