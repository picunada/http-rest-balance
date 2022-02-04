package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgconn"
	"github.com/picunada/http-rest-balance/internal/users"
	"github.com/picunada/http-rest-balance/pkg/logging"
	"github.com/picunada/http-rest-balance/pkg/postgresql"
)

type repository struct {
	client postgresql.Client
	logger *logging.Logger
}

func (r *repository) Create(ctx context.Context, user users.User) (string, error) {
	q := `
		INSERT INTO "user" (name)
		VALUES ($1)
		RETURNING id
		`
	if err := r.client.QueryRow(ctx, q, user.Name).Scan(&user.ID); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s", pgErr.Message)
		}
	}

}

func (r *repository) FindAll(ctx context.Context) (u []users.User, err error) {
	//TODO implement me
	panic("implement me")
}

func (r *repository) FindOne(ctx context.Context, id string) (users.User, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repository) Update(ctx context.Context, user users.User) error {
	//TODO implement me
	panic("implement me")
}

func (r *repository) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func Newrepository(client postgresql.Client, logger *logging.Logger) users.Repository {
	return &*repository{
		client: client,
		logger: logger,
	}
}
