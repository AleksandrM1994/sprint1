package middleware

import (
	"context"
	"net/http"

	"go.uber.org/zap"

	"github.com/sprint1/internal/app/shortener/service"
)

type contextKey string

const (
	UserID contextKey = "user_id"
)

func Authenticate(lg *zap.SugaredLogger, s service.Service, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("auth_cookie")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		userID, errCheckCookie := s.CheckCookie(cookie.Value)
		if errCheckCookie != nil {
			w.WriteHeader(http.StatusUnauthorized)
			http.Error(w, errCheckCookie.Error(), http.StatusUnauthorized)
			return
		}

		ctx := r.Context()
		newCtx := context.WithValue(ctx, UserID, userID)
		r = r.WithContext(newCtx)

		next.ServeHTTP(w, r)
	})
}
