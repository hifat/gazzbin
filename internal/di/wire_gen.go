// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/google/wire"
	"gazzbin/internal/config"
	"gazzbin/internal/handler"
	"gazzbin/internal/handler/taskHdl"
	"gazzbin/internal/repository"
	"gazzbin/internal/repository/taskRepo"
	"gazzbin/internal/service/taskSrv"
)

// Injectors from wire.go:

func InitializeAPI(config2 config.AppConfig) (Adapter, func()) {
	db, cleanup := repository.NewPostgresConnection(config2)
	taskRepository := taskRepo.NewTaskRepository(db)
	taskService := taskSrv.NewTaskService(taskRepository)
	taskHandler := taskHdl.NewTaskHandler(taskService)
	handlerHandler := handler.NewHandler(taskHandler)
	adapter := NewAdapter(handlerHandler)
	return adapter, func() {
		cleanup()
	}
}

// wire.go:

var RepoSet = wire.NewSet(repository.NewPostgresDBSet, taskRepo.TaskRepoSet)

var ServiceSet = wire.NewSet(taskSrv.TaskServiceSet)

var HandlerSet = wire.NewSet(taskHdl.TaskHandlerSet, handler.NewHandlerSet)
