package repository

import (
	"context"
	"go-shorts/internal/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UrlRepository struct {
	db *pgxpool.Pool
}

type Repository interface {
	Create(u *model.URL) error
	FindByShortCode(code string) (*model.URL, error)
	IncreaseClick(code string) error
}

func NewUrlRepository(db *pgxpool.Pool) *UrlRepository {
	return &UrlRepository{
		db: db,
	}
}

func (r *UrlRepository) Create(u *model.URL) error {
	q := `
		INSERT INTO urls (id, original_url, short_code, clicks_count, expires_at)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := r.db.Exec(
		context.Background(),
		q,
		u.OriginalURL,
		u.ShortCode,
		u.ClicksCount,
	)

	return err
}

func (r *UrlRepository) FindByShortCode(code string) (*model.URL, error) {
	q := NewQueryBuilder("urls").
		Select("short_code", "original_url", "clicks_count").
		Where("short_code = $1").
		Build()

	var url model.URL

	err := r.db.QueryRow(
		context.Background(),
		q,
		code,
	).Scan(
		&url.ShortCode,
		&url.OriginalURL,
		&url.ClicksCount,
	)

	if err != nil {
		return nil, err
	}

	return &url, nil
}

func (r *UrlRepository) IncreaseClick(code string) error {
	q := `
	UPDATE urls
	SET clicks_count = clicks_count + 1
	WHERE short_code = $1
	`
	_, err := r.db.Exec(
		context.Background(),
		q,
		code,
	)
	return err
}
