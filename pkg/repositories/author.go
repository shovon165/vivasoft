package repositories

import (
	"book-crud/pkg/domain"
	"book-crud/pkg/models"
	"sync"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type authorRepo struct {
	db *gorm.DB
}

func AuthorDBInstance(d *gorm.DB) domain.IAuthorRepo {
	return &authorRepo{
		db: d,
	}
}

func (repo *authorRepo) GetAuthors(request map[string]string) ([]models.AuthorDetail, error) {
	var authorDetails []models.AuthorDetail
	if err := repo.db.Find(&authorDetails).Error; err != nil {
		return nil, err
	}
	parsedSchema, err := schema.Parse(&models.AuthorDetail{}, &sync.Map{}, schema.NamingStrategy{})
	if err != nil {
		return nil, err
	}
	for key, value := range request {
		mappedFieldInDB := parsedSchema.FieldsByName[key].DBName
		err := repo.db.Where(mappedFieldInDB+" = ?", value).Find(&authorDetails).Error
		if err != nil {
			return nil, err
		}
	}

	return authorDetails, nil
}

func (repo *authorRepo) GetAuthor(authorID uint) (*models.AuthorDetail, error) {
	authorDetail := &models.AuthorDetail{}
	if err := repo.db.Where("id = ?", authorID).First(authorDetail).Error; err != nil {
		return nil, err
	}
	return authorDetail, nil
}

func (repo *authorRepo) CreateAuthor(author *models.AuthorDetail) error {
	if err := repo.db.Create(author).Error; err != nil {
		return err
	}
	return nil
}

func (repo *authorRepo) UpdateAuthor(author *models.AuthorDetail) error {
	if err := repo.db.Save(author).Error; err != nil {
		return err
	}
	return nil
}

func (repo *authorRepo) DeleteAuthor(authorID uint) error {
	authorDetail := &models.AuthorDetail{}
	if err := repo.db.Where("id = ?", authorID).Delete(authorDetail).Error; err != nil {
		return err
	}
	return nil
}
