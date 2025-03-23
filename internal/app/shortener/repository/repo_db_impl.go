package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/jackc/pgerrcode"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"go.uber.org/zap"

	"github.com/sprint1/config"
	custom_errs "github.com/sprint1/internal/app/shortener/errors"
	"github.com/sprint1/internal/app/shortener/helpers"
)

const (
	migrations       = "internal/app/shortener/repository/migrations/0001_create_database.sql"
	createURL        = "internal/app/shortener/repository/scripts/create_url.sql"
	getURLByShortURL = "internal/app/shortener/repository/scripts/get_url_by_short_url.sql"
	createUser       = "internal/app/shortener/repository/scripts/create_user.sql"
	getUser          = "internal/app/shortener/repository/scripts/get_user.sql"
	updateUser       = "internal/app/shortener/repository/scripts/update_user.sql"
	getUserByID      = "internal/app/shortener/repository/scripts/get_user_by_id.sql"
)

type RepoDBImpl struct {
	lg  *zap.SugaredLogger
	cfg config.Config
	db  *sqlx.DB
}

func NewRepoDBImpl(logger *zap.SugaredLogger, cfg config.Config) (*RepoDBImpl, error) {
	db, err := Connect(cfg.DNS)
	if err != nil {
		return nil, err
	}
	errMigrate := Migrate(db)
	if errMigrate != nil {
		return nil, errMigrate
	}
	return &RepoDBImpl{
		lg:  logger,
		cfg: cfg,
		db:  db,
	}, nil
}

func Connect(dns string) (*sqlx.DB, error) {
	db, errConnect := sqlx.Connect("postgres", dns)
	if errConnect != nil {
		return nil, fmt.Errorf("sqlx.Connect: %w", errConnect)
	}
	return db, nil
}

func Migrate(db *sqlx.DB) error {
	scripts, errReadFile := helpers.ReadFile(migrations)
	if errReadFile != nil {
		return fmt.Errorf("db.ReadFile: %w", errReadFile)
	}

	// Разделяем файл на отдельные SQL-команды.
	sqlCommands := strings.Split(scripts, ";")
	// Выполняем их в цикле.
	for _, command := range sqlCommands {
		switch command {
		case "", "\n":
		default:
			_, errExec := db.Exec(command)
			if errExec != nil {
				fmt.Printf("commandr: %v\n", command)
				return fmt.Errorf("db.Exec: %w", errExec)
			}
		}
	}
	return nil
}

func (r *RepoDBImpl) Ping() error {
	return r.db.Ping()
}

func (r *RepoDBImpl) CreateURL(shortURL, originalURL, userID string) error {
	script, errReadFile := helpers.ReadFile(createURL)
	if errReadFile != nil {
		return fmt.Errorf("db.ReadFile: %w", errReadFile)
	}

	_, errNamedExec := r.db.NamedExec(script, &URL{ShortURL: shortURL, OriginalURL: originalURL, UserID: userID})
	if errNamedExec != nil {
		var pgErr *pq.Error
		if errors.As(errNamedExec, &pgErr) {
			code := pgErr.Code
			switch code {
			case pgerrcode.UniqueViolation:
				return custom_errs.ErrUniqueViolation
			default:
				return fmt.Errorf("db.NamedExec: %w", errNamedExec)
			}
		}
		return fmt.Errorf("db.NamedExec: %w", errNamedExec)
	}

	return nil
}

func (r *RepoDBImpl) GetURLByShortURL(shortURL string) (*URL, error) {
	script, errReadFile := helpers.ReadFile(getURLByShortURL)
	if errReadFile != nil {
		return nil, fmt.Errorf("db.ReadFile: %w", errReadFile)
	}

	url := &URL{}
	errGet := r.db.Get(url, script, shortURL)
	if errGet != nil {
		if errors.Is(errGet, sql.ErrNoRows) {
			return nil, custom_errs.ErrNotFound
		}
		return nil, fmt.Errorf("db.Get: %w", errGet)
	}
	return url, nil
}

func (r *RepoDBImpl) CreateUser(id, login, password string) error {
	script, errReadFile := helpers.ReadFile(createUser)
	if errReadFile != nil {
		return fmt.Errorf("db.ReadFile: %w", errReadFile)
	}

	_, errNamedExec := r.db.NamedExec(script, &User{ID: id, Login: login, Password: password})
	if errNamedExec != nil {
		var pgErr *pq.Error
		if errors.As(errNamedExec, &pgErr) {
			code := pgErr.Code
			switch code {
			case pgerrcode.UniqueViolation:
				return custom_errs.ErrUniqueViolation
			default:
				return fmt.Errorf("db.NamedExec: %w", errNamedExec)
			}
		}
		return fmt.Errorf("db.NamedExec: %w", errNamedExec)
	}

	return nil
}

func (r *RepoDBImpl) GetUser(login, password string) (*User, error) {
	script, errReadFile := helpers.ReadFile(getUser)
	if errReadFile != nil {
		return nil, fmt.Errorf("db.ReadFile: %w", errReadFile)
	}

	user := &User{}
	errGet := r.db.Get(user, script, login, password)
	if errGet != nil {
		if errors.Is(errGet, sql.ErrNoRows) {
			return nil, custom_errs.ErrNotFound
		}
		return nil, fmt.Errorf("db.Get: %w", errGet)
	}
	return user, nil
}

func (r *RepoDBImpl) UpdateUser(user *User) error {
	script, errReadFile := helpers.ReadFile(updateUser)
	if errReadFile != nil {
		return fmt.Errorf("db.ReadFile: %w", errReadFile)
	}

	_, errNamedExec := r.db.NamedExec(script, user)
	if errNamedExec != nil {
		return fmt.Errorf("db.NamedExec: %w", errNamedExec)
	}
	return nil
}

func (r *RepoDBImpl) GetUserByID(id string) (*User, error) {
	script, errReadFile := helpers.ReadFile(getUserByID)
	if errReadFile != nil {
		return nil, fmt.Errorf("db.ReadFile: %w", errReadFile)
	}

	user := &User{}
	errGet := r.db.Get(user, script, id)
	if errGet != nil {
		if errors.Is(errGet, sql.ErrNoRows) {
			return nil, custom_errs.ErrNotFound
		}
		return nil, fmt.Errorf("db.Get: %w", errGet)
	}
	return user, nil
}
