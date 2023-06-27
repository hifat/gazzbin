package httpResponse

import (
	"gazzbin/internal/util/uError"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Error(ctx *gin.Context, err error) {
	if errRes, ok := err.(uError.ErrorResponse); ok {
		ctx.AbortWithStatusJSON(errRes.Status, errRes)
		return
	}

	ctx.AbortWithStatusJSON(http.StatusInternalServerError, uError.ErrorResponse{
		Message: http.StatusText(http.StatusInternalServerError),
	})
}

func BadRequest(ctx *gin.Context, err error) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest, uError.ErrorResponse{
		Message: err.Error(),
	})
}
