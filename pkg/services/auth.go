package services

import (
	"book-crud/pkg/domain"
	"book-crud/pkg/models"
	"book-crud/pkg/types"
	"book-crud/pkg/utils"
	"errors"
)

type authService struct {
	userRepo domain.IUserRepo
}

func AuthServiceInstance(userRepo domain.IUserRepo) domain.IAuthService {
	return &authService{
		userRepo: userRepo,
	}
}

func (service *authService) LoginUser(loginRequest *types.LoginRequest) (*types.LoginResponse, error) {
	existingUser, err := service.userRepo.GetUser(&loginRequest.UserName)
	if err != nil {
		return nil, errors.New("user does not exist")
	}

	if err := utils.ComparePassword(existingUser.PasswordHash, loginRequest.Password); err != nil {
		return nil, errors.New("incorrect password")
	}
	// Generate JWT token
	token, err := utils.GetJwtForUser(existingUser.Username)
	if err != nil {
		return nil, err
	}

	return &types.LoginResponse{
		Token: token,
	}, nil

}

func (service *authService) SignupUser(registerRequest *types.SignupRequest) error {
	passwordHash, err := utils.HashPassword(registerRequest.Password)
	if err != nil {
		return err
	}
	user := &models.UserDetail{
		Username:     registerRequest.UserName,
		PasswordHash: passwordHash,
		Name:         registerRequest.Name,
		Email:        registerRequest.Email,
		Address:      registerRequest.Address,
	}
	if err := service.userRepo.CreateUser(user); err != nil {
		return err
	}

	return nil
}
