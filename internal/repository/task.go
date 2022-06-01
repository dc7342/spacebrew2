package repository

import (
	"github.com/je09/spacebrew2/internal/entity"
	"gorm.io/gorm"
)

type TaskGORM struct {
	db *gorm.DB
}

func NewTaskGORM(db *gorm.DB) *TaskGORM {
	return &TaskGORM{db: db}
}

func (t *TaskGORM) GetAll() ([]entity.Task, error) {
	var tt []entity.Task
	res := t.db.Model(TaskModel{}).Find(&tt)
	if res.Error != nil {
		return nil, res.Error
	}

	return tt, nil
}

func (t *TaskGORM) Get(id int) (entity.Task, error) {
	var tt entity.Task
	res := t.db.Model(TaskModel{}).Find(tt).Where("id = ?", id)
	if res.Error != nil {
		return entity.Task{}, res.Error
	}

	return tt, nil
}

func (t *TaskGORM) Close(id int) error {
	var tt entity.Task
	res := t.db.Model(TaskModel{}).Find(tt).Where("id = ?", id).Update("open", false)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (t *TaskGORM) Open(id int) error {
	var tt entity.Task
	res := t.db.Model(TaskModel{}).Find(tt).Where("id = ?", id).Update("open", true)
	if res.Error != nil {
		return res.Error
	}

	return nil
}
