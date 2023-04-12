package dto

import "time"

type GameDTOInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type GameDTOOutput struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdat"`
	UpdatedAt   time.Time `json:"updatedat"`
}
