package taskRepo

import (
	"go-casbin/internal/domain/taskDomain"
	"go-casbin/internal/model/gormModel"

	"github.com/google/uuid"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var TaskRepoSet = wire.NewSet(NewTaskRepository)

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) taskDomain.TaskRepository {
	return &taskRepository{db}
}

func (r taskRepository) Get(userID *uuid.UUID, res *[]taskDomain.ResTask) (err error) {
	tx := r.db.Model(&gormModel.Task{})

	if userID != nil {
		tx.Where("user_id = ?", userID)
	}

	return tx.Find(res).Error
}

func (r taskRepository) GetByID(id uuid.UUID, userID *uuid.UUID, res *taskDomain.ResTask) (err error) {
	tx := r.db.Model(&gormModel.Task{}).
		Where("id = ?", id)

	if userID != nil {
		tx.Where("user_id = ?", userID)
	}

	if err = tx.First(&res).Error; err != nil {
		return err
	}

	return nil
}

func (r taskRepository) Create(req taskDomain.ReqTask) (res *taskDomain.TaskOG, err error) {
	return res, r.db.Create(&gormModel.Task{
		Name:   req.Name,
		Done:   req.Done,
		UserID: req.UserID,
	}).Scan(&res).Error
}

func (r taskRepository) Update(id uuid.UUID, req taskDomain.ReqTask) (res *taskDomain.TaskOG, err error) {
	return res, r.db.Model(&gormModel.Task{}).
		Where("id = ?", id).
		Updates(map[string]any{
			"name": req.Name,
			"done": req.Done,
		}).
		Scan(&res).Error
}

func (r taskRepository) Delete(id uuid.UUID) (err error) {
	return r.db.Where("id = ?", id).
		Delete(&gormModel.Task{}).Error
}
