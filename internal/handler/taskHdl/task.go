package taskHdl

import (
	"go-casbin/internal/domain/taskDomain"
	"go-casbin/internal/handler/httpResponse"
	"go-casbin/internal/util/uResponse"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/google/wire"
)

var TaskHandlerSet = wire.NewSet(NewTaskHandler)

type TaskHandler struct {
	taskSrv taskDomain.TaskService
}

func NewTaskHandler(taskSrv taskDomain.TaskService) TaskHandler {
	return TaskHandler{taskSrv}
}

func (h TaskHandler) Get(ctx *gin.Context) {
	res := []taskDomain.ResTask{}
	err := h.taskSrv.Get(nil, &res)
	if err != nil {
		httpResponse.Error(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, uResponse.SuccessResponse{
		Items: res,
	})
}

func (h TaskHandler) GetByID(ctx *gin.Context) {
	taskID, err := uuid.Parse(ctx.Param("taskID"))
	if err != nil {
		httpResponse.BadRequest(ctx, err)
		return
	}

	res := taskDomain.ResTask{}
	err = h.taskSrv.GetByID(taskID, nil, &res)
	if err != nil {
		httpResponse.Error(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, uResponse.SuccessResponse{
		Item: res,
	})
}

func (h TaskHandler) Create(ctx *gin.Context) {
	var req taskDomain.ReqTask
	err := ctx.ShouldBind(&req)
	if err != nil {
		httpResponse.BadRequest(ctx, err)
		return
	}

	res, err := h.taskSrv.Create(req)
	if err != nil {
		httpResponse.Error(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, uResponse.SuccessResponse{
		Item: res,
	})
}

func (h TaskHandler) Update(ctx *gin.Context) {
	taskID, err := uuid.Parse(ctx.Param("taskID"))
	if err != nil {
		httpResponse.BadRequest(ctx, err)
		return
	}

	var req taskDomain.ReqTask
	err = ctx.ShouldBind(&req)
	if err != nil {
		httpResponse.BadRequest(ctx, err)
		return
	}

	res, err := h.taskSrv.Update(taskID, nil, req)
	if err != nil {
		httpResponse.Error(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, uResponse.SuccessResponse{
		Item: res,
	})
}

func (h TaskHandler) Delete(ctx *gin.Context) {
	taskID, err := uuid.Parse(ctx.Param("taskID"))
	if err != nil {
		httpResponse.BadRequest(ctx, err)
		return
	}

	err = h.taskSrv.Delete(taskID, nil)
	if err != nil {
		httpResponse.Error(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, uResponse.SuccessResponse{
		Message: "deleted task",
	})
}
