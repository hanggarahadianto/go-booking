package models

import (
	"time"

	"github.com/google/uuid"
)

type Building struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name      string    `json:"name"`
	Content   string    `json:"content"`
	CreatedAt time.Time `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at,omitempty"`
}

// type CreatePostInput struct {
// 	Title     string    `json:"title"  binding:"required"`
// 	Content   string    `json:"content" binding:"required"`
// 	User      string    `json:"user,omitempty"`
// 	CreatedAt time.Time `json:"created_at,omitempty"`
// 	UpdatedAt time.Time `json:"updated_at,omitempty"`
// }

// type UpdatePost struct {
// 	Title     string    `json:"title,omitempty"`
// 	Content   string    `json:"content,omitempty"`
// 	User      string    `json:"user,omitempty"`
// 	CreateAt  time.Time `json:"created_at,omitempty"`
// 	UpdatedAt time.Time `json:"updated_at,omitempty"`
// }
