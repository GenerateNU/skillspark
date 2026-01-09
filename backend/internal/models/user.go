package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID                   uuid.UUID `json:"id" db:"id"`
	Name                 string    `json:"name" db:"name"`
	Email                string    `json:"email" db:"email"`
	Username             string    `json:"username" db:"username"`
	ProfilePictureS3Key  *string   `json:"profile_picture_s3_key" db:"profile_picture_s3_key"`
	CreatedAt            time.Time `json:"created_at" db:"created_at"`
	UpdatedAt            time.Time `json:"updated_at" db:"updated_at"`
}
