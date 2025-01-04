package usecases

import (
	"server/infrastructure/app"
	"server/internal/core/models"
	"server/internal/core/repositories"
)

type UserUsecase interface {
	GetUserById(id int, requesterUserId int) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
}

type userUsecase struct {
	userRepo repositories.UserRepository
}

func NewUserUsecase(userRepo repositories.UserRepository) UserUsecase {
	return &userUsecase{userRepo}
}

func (u *userUsecase) GetUserById(id int, requesterUserId int) (*models.User, error) {
	if requesterUserId != id {
		return nil, app.ErrUnauthorized
	}

	user, err := u.userRepo.GetUserById(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userUsecase) GetUserByEmail(email string) (*models.User, error) {
	user, err := u.userRepo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	return user, nil
}
