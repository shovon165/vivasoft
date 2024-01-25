package repositories

import (
	"book-crud/pkg/domain"
	"book-crud/pkg/models"
	"errors"

	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func UserDBInstance(d *gorm.DB) domain.IUserRepo {
	return &userRepo{
		db: d,
	}
}

func (repo *userRepo) GetUser(username *string) (*models.UserDetail, error) {
	user := &models.UserDetail{}
	if err := repo.db.Where("username = ?", username).First(user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (repo *userRepo) CreateUser(user *models.UserDetail) error {
	if err := repo.db.Create(user).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errors.New("username already exists")
		}
		return err
	}
	return nil
}
