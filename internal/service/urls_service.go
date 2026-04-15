// internal/service/url_service.go
package service

import (
	"go-shorts/internal/model"
	"go-shorts/internal/queue"
	"go-shorts/internal/repository"
	util "go-shorts/pkg/utils"

	"github.com/google/uuid"
)

type UrlService struct {
	repo *repository.UrlRepository
}

func NewUrlService(repo *repository.UrlRepository) *UrlService {
	return &UrlService{repo: repo}
}

func (s *UrlService) Create(url string, expires_at *string) (string, error) {
	id := uuid.New()
	shortCode := util.GenerateShortCodeFromURL(url)

	if error := s.repo.Create(&model.URL{
		ID:          id,
		ShortCode:   shortCode,
		OriginalURL: url,
		ClicksCount: 0,
		ExpiresAt:   expires_at,
	}); error != nil {
		return "", error
	}

	return shortCode, nil
}

func (s *UrlService) GetOriginalURL(code string) (string, error) {
	u, err := s.repo.FindByShortCode(code)
	if err != nil {
		return "", err
	}
	queue.NewClickQueue(s.repo)
	return u.OriginalURL, nil
}
