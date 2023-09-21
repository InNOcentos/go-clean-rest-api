package user

import (
	"context"
	"fmt"

	"github.com/InNOcentos/go-clean-rest-api/internal/entity"
	"github.com/InNOcentos/go-clean-rest-api/pkg/database"
)

type Repository interface {
	CreateUser(context.Context, *entity.User) (*entity.User, error)
	GetUser(context.Context, string) (*entity.User, error)
}

type repository struct {
	db *database.Postgres
}

func NewRepository(pg *database.Postgres) *repository {
	return &repository{
		db: pg,
	}
}

func (r *repository) CreateUser(ctx context.Context, u *entity.User) (*entity.User, error) {
  var userId string
  query := "INSERT INTO users (name) VALUES ($1) RETURNING id"
  fmt.Println(123)

  row := r.db.Pool.QueryRow(ctx, query, u.Name)
  if err := row.Scan(&userId); err != nil {
  fmt.Println(err)
    return nil, err
  }

  fmt.Println(123)

  u.Id = userId

	return u, nil
}

func (r *repository) GetUser(ctx context.Context, id string) (*entity.User, error) {
  query := "SELECT * from users WHERE id = $1 LIMIT 1"

  var user entity.User

  row := r.db.Pool.QueryRow(ctx, query, id)
  if err := row.Scan(&user.Id, &user.Name); err != nil {
  fmt.Println(123)
    fmt.Println(err)
    return nil, err
  }


	return &user, nil
}
