package repository

import (
	"context"
	"testing"

	"github.com/google/uuid"
)

var repo RepoDBImpl

func init() {
	db, _ := Connect("user=postgres password=postgres dbname=praktikum host=postgres port=5432 sslmode=disable")
	Migrate(db)
	repo = RepoDBImpl{db: db}
}

func BenchmarkRepository_Ping(b *testing.B) {
	ctx := context.Background()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = repo.Ping(ctx)
	}
}

func BenchmarkRepository_CreateURLs(b *testing.B) {
	ctx := context.Background()

	urls := []*URL{
		{
			ID:          1,
			ShortURL:    "996e1f714b08",
			OriginalURL: "https://github.com",
			UserID:      "39e529f3-7947-4d3b-aee4-d49a3a757c0f",
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = repo.CreateURLs(ctx, urls)
	}
}

func BenchmarkRepository_CreateUser(b *testing.B) {
	ctx := context.Background()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = repo.CreateUser(ctx, uuid.NewString(), uuid.NewString(), uuid.NewString())
	}
}

func BenchmarkRepository_GetUser(b *testing.B) {
	ctx := context.Background()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = repo.GetUser(ctx, uuid.NewString(), uuid.NewString())
	}
}

func BenchmarkRepository_UpdateUser(b *testing.B) {
	ctx := context.Background()

	user := &User{
		ID:       uuid.NewString(),
		Login:    uuid.NewString(),
		Password: uuid.NewString(),
		Cookie:   uuid.NewString(),
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = repo.UpdateUser(ctx, user)
	}
}

func BenchmarkRepository_GetUserByID(b *testing.B) {
	ctx := context.Background()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = repo.GetUserByID(ctx, uuid.NewString())
	}
}

func BenchmarkRepository_GetURLsByUserID(b *testing.B) {
	ctx := context.Background()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = repo.GetURLsByUserID(ctx, uuid.NewString())
	}
}

func BenchmarkRepository_MakeURLsDeleted(b *testing.B) {
	ctx := context.Background()

	urls := []*URL{
		{
			ID:          1,
			ShortURL:    "996e1f714b08",
			OriginalURL: "https://github.com",
			UserID:      "39e529f3-7947-4d3b-aee4-d49a3a757c0f",
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = repo.MakeURLsDeleted(ctx, urls)
	}
}
