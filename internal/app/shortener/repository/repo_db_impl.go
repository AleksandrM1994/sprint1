package repository

import (
	"context"
	"database/sql"
	"embed"
	"errors"
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

func (r *RepoDBImpl) CreateURL(ctx context.Context, shortURL, originalURL, userID string) error {
	_, errExecContext := r.db.ExecContext(ctx, CreateURL, shortURL, originalURL, userID)
	if errExecContext == nil {
		return nil
	}
	var pgErr *pq.Error
	if ok := errors.As(errExecContext, &pgErr); ok {
		if pgErr.Code == pgerrcode.UniqueViolation {
			return custom_errs.ErrUniqueViolation
		}
	}
	return fmt.Errorf("db.ExecContext:%w", errExecContext)
}

func (r *RepoDBImpl) GetURLByShortURL(ctx context.Context, shortURL string) (*URL, error) {
	url := &URL{}
	errGet := r.db.GetContext(ctx, url, GetURLByShortURL, shortURL)
	if errGet == nil {
		return url, nil
	}
	if errors.Is(errGet, sql.ErrNoRows) {
		return nil, custom_errs.ErrNotFound
	}
	return nil, fmt.Errorf("db.GetContext: %w", errGet)
}

func (r *RepoDBImpl) CreateURLs(ctx context.Context, urls []*URL) error {
	tx, errBeginx := r.db.BeginTx(ctx, nil)
	if errBeginx != nil {
		return fmt.Errorf("error beginx: %w", errBeginx)
	}

	for _, url := range urls {
		_, errExecContext := tx.ExecContext(ctx, CreateURL, url.ShortURL, url.OriginalURL, url.UserID)
		if errExecContext != nil {
			if errRollback := tx.Rollback(); errRollback != nil {
				return fmt.Errorf("error rollback: %v", errRollback)
			}
			var pgErr *pq.Error
			if ok := errors.As(errExecContext, &pgErr); ok {
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

func (r *RepoDBImpl) CreateUser(ctx context.Context, id, login, password string) error {
	_, errExecContext := r.db.ExecContext(ctx, createUser, id, login, password)
	if errExecContext == nil {
		return nil
	}
	var pgErr *pq.Error
	if ok := errors.As(errExecContext, &pgErr); ok {
		if pgErr.Code == pgerrcode.UniqueViolation {
			return custom_errs.ErrUniqueViolation
		}
	}
	return fmt.Errorf("db.ExecContext:%w", errExecContext)
}

func (r *RepoDBImpl) GetUser(ctx context.Context, login, password string) (*User, error) {
	user := &User{}
	errGet := r.db.GetContext(ctx, user, getUser, login, password)
	if errGet == nil {
		return user, nil
	}
	if errors.Is(errGet, sql.ErrNoRows) {
		return nil, custom_errs.ErrNotFound
	}
	return nil, fmt.Errorf("db.GetContext: %w", errGet)
}

func (r *RepoDBImpl) UpdateUser(ctx context.Context, user *User) error {
	_, errExecContext := r.db.ExecContext(ctx, updateUser, user.Cookie, user.CookieFinish, user.ID)
	if errExecContext == nil {
		return nil
	}
	var pgErr *pq.Error
	if ok := errors.As(errExecContext, &pgErr); ok {
		if pgErr.Code == pgerrcode.UniqueViolation {
			return custom_errs.ErrUniqueViolation
		}
	}
	return fmt.Errorf("db.ExecContext:%w", errExecContext)
}

func (r *RepoDBImpl) GetUserByID(ctx context.Context, id string) (*User, error) {
	user := &User{}
	errGet := r.db.GetContext(ctx, user, getUserByID, id)
	if errGet == nil {
		return user, nil
	}
	if errors.Is(errGet, sql.ErrNoRows) {
		return nil, custom_errs.ErrNotFound
	}
	return nil, fmt.Errorf("db.GetContext: %w", errGet)
}

func (r *RepoDBImpl) GetURLsByUserID(ctx context.Context, id string) ([]*URL, error) {
	var urls []*URL
	errSelectContext := r.db.SelectContext(ctx, &urls, GetURLsByUserID, id)
	if errSelectContext == nil {
		return urls, nil
	}
	if errors.Is(errSelectContext, sql.ErrNoRows) {
		return nil, custom_errs.ErrNoContent
	}
	return nil, fmt.Errorf("db.SelectContext: %w", errSelectContext)
}
