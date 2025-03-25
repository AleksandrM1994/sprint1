package repository

import (
	"context"
	"database/sql"
	"embed"
	"fmt"

	"github.com/jackc/pgerrcode"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"go.uber.org/zap"

	"github.com/sprint1/config"
	custom_errs "github.com/sprint1/internal/app/shortener/errors"
)

type RepoDBImpl struct {
	lg  *zap.SugaredLogger
	cfg config.Config
	db  *sqlx.DB
}

func Connect(dns string) (*sqlx.DB, error) {
	db, errConnect := sqlx.Connect("postgres", dns)
	if errConnect != nil {
		return nil, fmt.Errorf("sqlx.Connect: %w", errConnect)
	}
	return db, nil
}

//go:embed migrations/*.sql
var embedMigrations embed.FS

func Migrate(db *sqlx.DB) error {
	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("error setting SQL dialect: %w", err)
	}

	if err := goose.Up(db.DB, "migrations"); err != nil {
		return fmt.Errorf("error migration: %w", err)
	}
	return nil
}

func (r *RepoDBImpl) Ping(ctx context.Context) error {
	return r.db.PingContext(ctx)
}

func (r *RepoDBImpl) CreateURL(ctx context.Context, shortURL string, originalURL string) error {
	_, errExecContext := r.db.ExecContext(ctx, CreateURL, shortURL, originalURL)
	if errExecContext == nil {
		return nil
	}
	pgErr, ok := errExecContext.(*pq.Error)
	if ok {
		if pgErr.Code == pgerrcode.UniqueViolation {
			return custom_errs.ErrUniqueViolation
		}
	}
	return fmt.Errorf("db.NamedExec:%w", errExecContext)
}

func (r *RepoDBImpl) GetURLByShortURL(ctx context.Context, shortURL string) (*URL, error) {
	url := &URL{}
	errGet := r.db.GetContext(ctx, url, GetURLByShortURL, shortURL)
	if errGet == nil {
		return url, nil
	}
	if errGet == sql.ErrNoRows {
		return nil, custom_errs.ErrNotFound
	}
	return nil, fmt.Errorf("db.Get: %w", errGet)
}

func (r *RepoDBImpl) CreateURLs(ctx context.Context, urls []*URL) error {
	tx, errBeginx := r.db.BeginTx(ctx, nil)
	if errBeginx != nil {
		return fmt.Errorf("error beginx: %w", errBeginx)
	}

	for _, url := range urls {
		_, errExecContext := tx.ExecContext(ctx, CreateURL, url.ShortURL, url.OriginalURL)
		if errExecContext != nil {
			if errRollback := tx.Rollback(); errRollback != nil {
				return fmt.Errorf("error rollback: %v", errRollback)
			}
			pgErr, ok := errExecContext.(*pq.Error)
			if ok {
				if pgErr.Code == pgerrcode.UniqueViolation {
					return custom_errs.ErrUniqueViolation
				}
			}
			return fmt.Errorf("error creating url: %w", errExecContext)
		}
	}

	if errCommit := tx.Commit(); errCommit != nil {
		return fmt.Errorf("error commit: %w", errCommit)
	}

	return nil
}
