package gormModel

import (
	"time"

	uuid "github.com/google/uuid"
)

type User struct {
	ID   uuid.UUID `gorm:"primaryKey; type:uuid; default:uuid_generate_v4()"`
	Name string    `gorm:"type:varchar(50); unique"`

	Task []Task

	CreatedAt *time.Time `gorm:"timestamp"`
	UpdatedAt *time.Time `gorm:"timestamp"`
	DeletedAt *time.Time `gorm:"timestamp"`
}
