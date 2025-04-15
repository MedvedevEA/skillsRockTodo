package postgresql

import (
	"context"
	"fmt"
	"log"
	"log/slog"

	"skillsRockTodo/internal/config"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgreSql struct {
	pool *pgxpool.Pool
	lg   *slog.Logger
}

func MustNew(ctx context.Context, lg *slog.Logger, cfg *config.PostgreSQL) *PostgreSql {
	const op = "postgresql.New"
	connString := fmt.Sprintf(
		`user=%s password=%s host=%s port=%d dbname=%s sslmode=%s pool_max_conns=%d pool_max_conn_lifetime=%s pool_max_conn_idle_time=%s`,
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
		cfg.SSLMode,
		cfg.PoolMaxConns,
		cfg.PoolMaxConnLifetime.String(),
		cfg.PoolMaxConnIdleTime.String(),
	)

	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		log.Fatalf("%s: %s", op, err)
	}

	config.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeCacheDescribe

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		log.Fatalf("%s: %s", op, err)
	}

	return &PostgreSql{
		pool,
		lg,
	}
}
