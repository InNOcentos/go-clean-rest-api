package user

import (
	"context"

	"github.com/InNOcentos/go-clean-rest-api/internal/entity"
	user "github.com/InNOcentos/go-clean-rest-api/internal/user/repository/postgres"
)

type UseCase interface {
	CreateUser(context.Context, CreateUserRequest) (*entity.User, error)
	GetUser(context.Context, string) (*entity.User, error)
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

func (u *useCase) CreateUser(ctx context.Context, input CreateUserRequest) (*entity.User, error) {
	user := &entity.User{
		Name: input.Name,
	}

	user, err := u.userRepo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *useCase) GetUser(ctx context.Context, id string) (*entity.User, error) {
	user, err := u.userRepo.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
