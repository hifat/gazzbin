package handler

import (
	"gazzbin/internal/handler/taskHdl"

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
