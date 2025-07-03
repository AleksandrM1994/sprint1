package interceptor

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// contextKey тип для хранения данных в контексте
type contextKey string

// UserID - идентификатор пользователя, который будет сохранен в контекст по итогу успешной авторизации
const (
	UserID contextKey = "user_id"
)

func AuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	var userID string
	// забираем userID из метаданных
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		values := md.Get("user_id")
		if len(values) > 0 {
			userID = values[0]
		}
	}

	// Добавляем userID в контекст
	ctxWithUserID := context.WithValue(ctx, UserID, userID)

	return handler(ctxWithUserID, req)
}
