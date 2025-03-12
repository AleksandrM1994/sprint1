package repository

type Repo interface {
	Ping() error
}
