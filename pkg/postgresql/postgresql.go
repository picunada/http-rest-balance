package postgresql

import (
	"context"
	"fmt"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/picunada/http-rest-balance/internal/config"
	"github.com/picunada/http-rest-balance/pkg/utils"
	"log"
	"time"
)

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{})
	Begin(ctx context.Context) (pgx.Tx, error)
}

func NewClient(ctx context.Context) (pool *pgxpool.Pool, err error) {
	cfg := config.GetConfig()
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", cfg.Storage.User, cfg.Storage.Password,
		cfg.Storage.Host, cfg.Storage.Port, cfg.Storage.Name)

	err = utils.DoWithTries(func() error {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		pool, err = pgxpool.Connect(ctx, dsn)
		if err != nil {
			fmt.Print("Failed to connect to PostgreSQL")
			return err
		}

		if err != nil {
			log.Fatal("Error while attempting to connect PostgreSQL")
		}

		return nil
	}, cfg.Storage.MaxAttempts, 5*time.Second)

	return pool, nil
}
