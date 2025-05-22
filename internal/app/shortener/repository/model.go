package repository

import "time"

// URL репозиторная структура урла
type URL struct {
	ID          int64  `db:"id"`
	ShortURL    string `db:"short_url"`
	OriginalURL string `db:"original_url"`
	UserID      string `db:"user_id"`
	IsDeleted   bool   `db:"is_deleted"`
}

// User репозиторная структура пользователя
type User struct {
	ID           string     `db:"id"`
	Login        string     `db:"login"`
	Password     string     `db:"password"`
	Cookie       string     `db:"cookie"`
	CookieFinish *time.Time `db:"cookie_finish"`
}
