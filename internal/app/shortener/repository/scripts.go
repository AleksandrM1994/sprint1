package repository

const (
	CreateURL        = "insert into urls (short_url, original_url) values ($1, $2)"
	GetURLByShortURL = "select id, short_url, original_url from urls where short_url=$1"
	createUser       = "insert into users (id, login, password) values ($1, $2, $3)"
	getUser          = "select id, login, password from users where login = $1 and password = $2"
	updateUser       = "update users set cookie = $1, cookie_finish = $2 where id = $3"
	getUserByID      = "select id, login, password, cookie, cookie_finish from users where id = $1"
	GetURLsByUserID  = "select id, short_url, original_url, user_id from urls where user_id = $1"
)
