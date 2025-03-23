package repository

import (
	"sync"

	custom_errs "github.com/sprint1/internal/app/shortener/errors"
)

type RepoMemoryImpl struct {
	urlStorage map[string]string
	mu         sync.Mutex
}

func NewRepoMemoryImpl() *RepoMemoryImpl {
	return &RepoMemoryImpl{
		urlStorage: make(map[string]string),
	}
}

func (r *RepoMemoryImpl) CreateURL(shortURL, originalURL, userID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, found := r.urlStorage[shortURL]; found {
		return custom_errs.ErrUniqueViolation
	}
	r.urlStorage[shortURL] = originalURL

	return nil
}

func (r *RepoMemoryImpl) GetURLByShortURL(shortURL string) (*URL, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	shortURLs := make([]string, len(r.urlStorage))
	for url := range r.urlStorage {
		shortURLs = append(shortURLs, url)
	}

	for _, k := range shortURLs {
		if k == shortURL {
			findURL := &URL{
				ID:          int64(len(shortURLs)),
				ShortURL:    k,
				OriginalURL: r.urlStorage[k],
			}
			return findURL, nil
		}
	}

	return nil, custom_errs.ErrNotFound
}

func (r *RepoMemoryImpl) Ping() error {
	return nil
}

func (r *RepoMemoryImpl) CreateUser(id, login, password string) error {
	return nil
}

func (r *RepoMemoryImpl) GetUser(login, password string) (*User, error) {
	return nil, nil
}

func (r *RepoMemoryImpl) UpdateUser(user *User) error {
	return nil
}
func (r *RepoMemoryImpl) GetUserByID(id string) (*User, error) {
	return nil, nil
}
