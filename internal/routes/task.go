package routes

import (
	"gazzbin/internal/handler/taskHdl"
)

func (r *routes) Task(h taskHdl.TaskHandler) {
	todo := r.router.Group("/tasks")
	{
		todo.GET("", h.Get)
		todo.POST("", h.Create)
		todo.GET("/:taskID", h.GetByID)
		todo.PATCH("/:taskID", h.Update)
		todo.DELETE("/:taskID", h.Delete)
	}
}
