package user

import (
	"log"

	"github.com/InNOcentos/go-clean-rest-api/internal/entity"
	"github.com/InNOcentos/go-clean-rest-api/pkg/database"
	"github.com/gofrs/uuid"
)

type Repository interface {
	CreateUser(user *entity.User) (*entity.User, error)
}

type repository struct {
	db *database.Postgres
}

func NewRepository(pg *database.Postgres) *repository {
	return &repository{
		db: pg,
	}
}

func (r *repository) CreateUser(u *entity.User) (*entity.User, error) {
	userId, err := uuid.NewV4()
	if err != nil {
		log.Panic(err)
	}

	u.Id = userId

	return u, nil
}
