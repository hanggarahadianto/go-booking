package models

import (
	"time"

	"github.com/google/uuid"
)

type Building struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name      string    `json:"name"`
	Content   string    `json:"content"`
	ImageURL  string    `json:"image_url"`
	CreatedAt time.Time `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at,omitempty"`
}

type CreateBuildingInput struct {
	Name     string `json:"name"  binding:"required"`
	Content  string `json:"content" binding:"required"`
	ImageURL string `json:"image_url"`
	// User      string    `json:"user,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
