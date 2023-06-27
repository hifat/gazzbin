//go:build wireinject
// +build wireinject

package di

import (
	"gazzbin/internal/config"
	"gazzbin/internal/handler"
	"gazzbin/internal/handler/taskHdl"
	"gazzbin/internal/repository"
	"gazzbin/internal/repository/taskRepo"
	"gazzbin/internal/service/taskSrv"

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
