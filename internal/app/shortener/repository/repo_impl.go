package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"

	"github.com/sprint1/config"
	"github.com/sprint1/internal/app/shortener/helpers"
)

const (
	URLsCreateTable  = "internal/app/shortener/repository/scripts/urls_create_table.sql"
	CreateURL        = "internal/app/shortener/repository/scripts/create_url.sql"
	GetURLByShortURL = "internal/app/shortener/repository/scripts/get_url_by_short_url.sql"
)

type RepoImpl struct {
	lg  *zap.SugaredLogger
	cfg config.Config
	db  *sqlx.DB
}

func NewRepoImpl(logger *zap.SugaredLogger, cfg config.Config) (*RepoImpl, error) {
	db, err := Connect(cfg.DNS)
	if err != nil {
		return nil, err
	}
	errMigrate := Migrate(db)
	if errMigrate != nil {
		return nil, errMigrate
	}
	return &RepoImpl{
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
	script, errReadFile := helpers.ReadFile(URLsCreateTable)
	if errReadFile != nil {
		return fmt.Errorf("db.ReadFile: %w", errReadFile)
	}
	_, errExec := db.Exec(script)
	if errExec != nil {
		return fmt.Errorf("db.Exec: %w", errExec)
	}
	return nil
}

func (r *RepoImpl) Ping() error {
	return r.db.Ping()
}

func (r *RepoImpl) CreateURL(shortURL string, originalURL string) error {
	script, errReadFile := helpers.ReadFile(CreateURL)
	if errReadFile != nil {
		return fmt.Errorf("db.ReadFile: %w", errReadFile)
	}

	_, errNamedExec := r.db.NamedExec(script, &URL{ShortURL: shortURL, OriginalURL: originalURL})
	if errNamedExec != nil {
		return fmt.Errorf("db.NamedExec: %w", errNamedExec)
	}

	return nil
}

func (r *RepoImpl) GetURLByShortURL(shortURL string) (*URL, error) {
	script, errReadFile := helpers.ReadFile(GetURLByShortURL)
	if errReadFile != nil {
		return nil, fmt.Errorf("db.ReadFile: %w", errReadFile)
	}

	url := &URL{}
	errGet := r.db.Get(url, script, shortURL)
	if errGet != nil {
		return nil, fmt.Errorf("db.Get: %w", errGet)

	}
	return url, nil
}
