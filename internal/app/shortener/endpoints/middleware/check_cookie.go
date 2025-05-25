package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"go.uber.org/zap"

	custom_errs "github.com/sprint1/internal/app/shortener/errors"
	"github.com/sprint1/internal/app/shortener/service"
)

// contextKey тип для хранения данных в контексте
type contextKey string

// UserID - идентификатор пользователя, который будет сохранен в контекст по итогу успешной авторизации
const (
	UserID contextKey = "user_id"
)

// Authenticate функция, отвечающая за авторизацию запросов
func Authenticate(lg *zap.SugaredLogger, s *service.ServiceImpl, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		cookies := r.Cookies()
		for _, cookie := range cookies {
			fmt.Printf("Name: %s, Value: %s\n", cookie.Name, cookie.Value)
		}

		cookie, err := r.Cookie("auth_cookie")
		if cookie != nil && err == nil && cookie.Value != "" {
			userID, errCheckCookie := s.CheckCookie(ctx, cookie.Value)
			if errCheckCookie != nil && !errors.Is(errCheckCookie, custom_errs.ErrNotFound) {
				w.WriteHeader(http.StatusInternalServerError)
				http.Error(w, errCheckCookie.Error(), http.StatusInternalServerError)
				return
			}

			if userID != "" {
				newCtx := context.WithValue(ctx, UserID, userID)
				r = r.WithContext(newCtx)
			} else {
				w.WriteHeader(http.StatusUnauthorized)
				http.Error(w, errors.New("empty user id").Error(), http.StatusUnauthorized)
				return
			}
		} else {
			login := uuid.New().String()
			psw, errCreateUser := s.CreateUser(ctx, login)
			if errCreateUser != nil {
				w.WriteHeader(http.StatusInternalServerError)
				http.Error(w, errCreateUser.Error(), http.StatusInternalServerError)
				return
			}

			user, errAuthenticateUser := s.AuthenticateUser(ctx, login, psw)
			if errAuthenticateUser != nil {
				w.WriteHeader(http.StatusInternalServerError)
				http.Error(w, errAuthenticateUser.Error(), http.StatusInternalServerError)
				return
			}

			http.SetCookie(w, &http.Cookie{
				Name:    "auth_cookie",
				Value:   user.Cookie,
				Path:    "/",
				Secure:  false,
				Expires: *user.CookieFinish,
			})

			newCtx := context.WithValue(ctx, UserID, user.ID)
			r = r.WithContext(newCtx)
		}

		next.ServeHTTP(w, r)
	})
}
