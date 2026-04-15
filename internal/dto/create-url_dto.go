package dto

type URLCreateDTO struct {
	URL       string  `json:"url"`
	ExpiresAt *string `json:"expires_at,omitempty"`
}
