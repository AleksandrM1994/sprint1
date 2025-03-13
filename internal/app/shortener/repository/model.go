package repository

type URL struct {
	Id          int64  `db:"id"`
	ShortURL    string `db:"short_url"`
	OriginalURL string `db:"original_url"`
}
