package handler

import (
	"go-casbin/internal/handler/taskHdl"

	"github.com/google/wire"
)

var NewHandlerSet = wire.NewSet(NewHandler)

type Handler struct {
	TaskHandler taskHdl.TaskHandler
}

func NewHandler(TaskHandler taskHdl.TaskHandler) Handler {
	return Handler{
		TaskHandler,
	}
}
