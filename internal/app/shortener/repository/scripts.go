package repository

// запросы к БД
const (
	// запрос по созданию пользователя
	CreateURL = "insert into urls (short_url, original_url, user_id) values ($1, $2, $3)"
	// запрос, который вернет полный урл по сокращенному
	GetURLByShortURL = "select id, short_url, original_url, user_id, is_deleted from urls where short_url=$1"
	// запрос по созданию пользователя
	createUser = "insert into users (id, login, password) values ($1, $2, $3)"
	// запрос, который вернет юзера по логину/паролю
	getUser = "select id, login, password from users where login = $1 and password = $2"
	// запрос по обновлению юзера по его айди
	updateUser = "update users set cookie = $1, cookie_finish = $2 where id = $3"
	// запрос по получению юзера по его айди
	getUserByID = "select id, login, password, cookie, cookie_finish from users where id = $1"
	// запрос по урлов по айди юзера
	GetURLsByUserID = "select id, short_url, original_url, user_id, is_deleted from urls where user_id = $1"
	// запрос по удалению урлов по их соращенной версии
	MakeURLDeleted = "update urls set is_deleted = $1 where short_url = $2"
	// запрос по получению количества пользователей
	GetUsersCount = "select count(*) from users"
	// запрос по получению количества сокреащенных урлов
	GetURLsCount = "select count(*) from urls"
)
