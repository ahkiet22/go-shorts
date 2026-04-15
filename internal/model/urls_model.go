package model

import "github.com/google/uuid"

type URL struct {
	ID          uuid.UUID `json:"id" db:"id"`
	OriginalURL string    `json:"original_url"`
	ShortCode   string    `json:"short_code"`

	ClicksCount int `json:"clicks_count"`

	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
	ExpiresAt *string `json:"expires_at,omitempty"`

	DeletedAt *string `json:"deleted_at,omitempty"`
}
