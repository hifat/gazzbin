//go:build wireinject
// +build wireinject

package di

import (
	"go-casbin/internal/config"
	"go-casbin/internal/handler"
	"go-casbin/internal/handler/taskHdl"
	"go-casbin/internal/repository"
	"go-casbin/internal/repository/taskRepo"
	"go-casbin/internal/service/taskSrv"

	"github.com/google/wire"
)

var RepoSet = wire.NewSet(
	repository.NewPostgresDBSet,
	taskRepo.TaskRepoSet,
)

var ServiceSet = wire.NewSet(
	taskSrv.TaskServiceSet,
)

var HandlerSet = wire.NewSet(
	taskHdl.TaskHandlerSet,
	handler.NewHandlerSet,
)

func InitializeAPI(config config.AppConfig) (Adapter, func()) {
	wire.Build(AdapterSet, RepoSet, ServiceSet, HandlerSet)
	return Adapter{}, nil
}
