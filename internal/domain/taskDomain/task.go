package taskDomain

import (
	"gazzbin/internal/domain/userDomain"
	"time"

	"github.com/google/uuid"
)

type TaskRepository interface {
	Get(userID *uuid.UUID, res *[]ResTask) (err error)
	GetByID(id uuid.UUID, userID *uuid.UUID, res *ResTask) (err error)
	Create(req ReqTask) (res *TaskOG, err error)
	Update(id uuid.UUID, req ReqTask) (res *TaskOG, err error)
	Delete(id uuid.UUID) (err error)
}

type TaskService interface {
	Get(userID *uuid.UUID, res *[]ResTask) (err error)
	GetByID(id uuid.UUID, userID *uuid.UUID, res *ResTask) (err error)
	Create(req ReqTask) (res *TaskOG, err error)
	Update(id uuid.UUID, userID *uuid.UUID, req ReqTask) (res *TaskOG, err error)
	Delete(id uuid.UUID, userID *uuid.UUID) (err error)
}

type TaskOG struct {
	ID        uuid.UUID  `json:"id"`
	Name      string     `json:"name"`
	Done      bool       `json:"done"`
	UserID    *uuid.UUID `json:"userID"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
}

type ResTask struct {
	TaskOG
	User *userDomain.UserOG `json:"user"`
}

type ReqTask struct {
	Name   string    `json:"name" binding:"required,max=50"`
	Done   bool      `json:"done,omitempty"`
	UserID uuid.UUID `json:"-"`
}
