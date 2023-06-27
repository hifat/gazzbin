package routes

import (
	"gazzbin/internal/handler"

	"github.com/gin-gonic/gin"
)

type routes struct {
	router  *gin.RouterGroup
	handler handler.Handler
}

func New(router *gin.RouterGroup, h handler.Handler) *routes {
	return &routes{
		router:  router,
		handler: h,
	}
}

func (r routes) Register() {
	r.Task(r.handler.TaskHandler)
}
