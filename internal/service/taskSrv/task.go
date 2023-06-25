package taskSrv

import (
	"go-casbin/internal/domain/taskDomain"
	"go-casbin/internal/util/uError"
	"log"

	"github.com/google/uuid"
	"github.com/google/wire"
)

var TaskServiceSet = wire.NewSet(NewTaskService)

type todoService struct {
	taskRepo taskDomain.TaskRepository
}

func NewTaskService(taskRepo taskDomain.TaskRepository) taskDomain.TaskService {
	return todoService{taskRepo}
}

func (s todoService) Get(userID *uuid.UUID, res *[]taskDomain.ResTask) (err error) {
	return s.taskRepo.Get(userID, res)
}

func (s todoService) GetByID(id uuid.UUID, userID *uuid.UUID, res *taskDomain.ResTask) (err error) {
	err = s.taskRepo.GetByID(id, userID, res)
	if err != nil {
		log.Println(err.Error())
		if err.Error() == uError.Message.RECORD_NOT_FOUND {
			return uError.RecordNotFound()
		}

		return uError.InternalError()
	}

	return nil
}

func (s todoService) Create(req taskDomain.ReqTask) (res *taskDomain.TaskOG, err error) {
	return s.taskRepo.Create(req)
}

func (s todoService) Update(id uuid.UUID, userID *uuid.UUID, req taskDomain.ReqTask) (res *taskDomain.TaskOG, err error) {
	err = s.taskRepo.GetByID(id, userID, nil)
	if err != nil {
		log.Println(err.Error())
		if err.Error() == uError.Message.RECORD_NOT_FOUND {
			return nil, uError.RecordNotFound()
		}

		return nil, uError.InternalError()
	}

	res, err = s.taskRepo.Update(id, req)
	if err != nil {
		return nil, uError.InternalError()
	}

	return res, nil
}

func (s todoService) Delete(id uuid.UUID, userID *uuid.UUID) (err error) {
	err = s.taskRepo.GetByID(id, userID, nil)
	if err != nil {
		log.Println(err.Error())
		if err.Error() == uError.Message.RECORD_NOT_FOUND {
			return uError.RecordNotFound()
		}

		return uError.InternalError()
	}

	return s.taskRepo.Delete(id)
}
