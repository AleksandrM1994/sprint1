package repository

import "time"

type URL struct {
	ID          int64  `db:"id"`
	ShortURL    string `db:"short_url"`
	OriginalURL string `db:"original_url"`
	UserID      string `db:"user_id"`
}

type User struct {
	ID           string     `db:"id"`
	Login        string     `db:"login"`
	Password     string     `db:"password"`
	Cookie       string     `db:"cookie"`
	CookieFinish *time.Time `db:"cookie_finish"`
}
