package repository

type Repo interface {
	Ping() error
	CreateURL(shortURL, originalURL, userID string) error
	GetURLByShortURL(shortURL string) (*URL, error)
	CreateUser(id, login, password string) error
	GetUser(login, password string) (*User, error)
	UpdateUser(user *User) error
	GetUserByID(id string) (*User, error)
}
