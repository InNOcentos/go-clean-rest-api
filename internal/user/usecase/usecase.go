package user

import (
	"github.com/InNOcentos/go-clean-rest-api/internal/entity"
	user "github.com/InNOcentos/go-clean-rest-api/internal/user/repository/postgres"
)

type UseCase interface {
	CreateUser(input CreateUserRequest) (*entity.User, error)
}

type useCase struct {
	userRepo user.Repository
}

func NewUseCase(userRepo user.Repository) *useCase {
	return &useCase{
		userRepo: userRepo,
	}
}

type CreateUserRequest struct {
	Name string `json:"name" binding:"required"`
}

func (u *useCase) CreateUser(input CreateUserRequest) (*entity.User, error) {
	user := &entity.User{
		Name: input.Name,
	}

	user, err := u.userRepo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
