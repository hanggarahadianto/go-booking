package models

import (
	"time"

	"github.com/google/uuid"
)

type Building struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name      string    `json:"name" form:"name"`
	Content   string    `json:"content" form:"content"`
	Image     string    `json:"image" form:"image"`
	CreatedAt time.Time `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at,omitempty"`
}

type CreateBuildingInput struct {
	Name    string `json:"name" form:"name"`
	Content string `json:"content" form:"content"`
	Image   string `json:"image" form:"image"`
	// User      string    `json:"user,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
