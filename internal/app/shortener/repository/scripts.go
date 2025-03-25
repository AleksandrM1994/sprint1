package repository

const (
	CreateURL        = "insert into urls (short_url, original_url) values ($1, $2);"
	GetURLByShortURL = "select id, short_url, original_url from urls where short_url=$1"
)
