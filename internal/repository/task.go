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

func (t *TaskGORM) GetByPage(page int, perPage int) ([]entity.Task, error) {
	var tt []entity.Task
	res := t.db.Model(TaskModel{}).Find(&tt).Where("id > ?", perPage*page).Limit(perPage)
	if res.Error != nil {
		return nil, res.Error
	}

	return tt, nil
}

func (t *TaskGORM) Count() (int64, error) {
	var c int64
	res := t.db.Model(TaskModel{}).Count(&c)
	if res.Error != nil {
		return 0, res.Error
	}

	return c, nil
}

func (t *TaskGORM) Get(id int64) (entity.Task, error) {
	var tt entity.Task
	res := t.db.Model(TaskModel{}).Find(tt).Where("id = ?", id)
	if res.Error != nil {
		return entity.Task{}, res.Error
	}

	return tt, nil
}

func (t *TaskGORM) Add(task entity.Task) error {
	res := t.db.Model(TaskModel{}).Create(task)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (t *TaskGORM) Update(task entity.Task) error {
	res := t.db.Model(TaskModel{}).Updates(task)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (t *TaskGORM) Close(id int64) error {
	var tt entity.Task
	res := t.db.Model(TaskModel{}).Find(tt).Where("id = ?", id).Update("open", false)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (t *TaskGORM) Open(id int64) error {
	var tt entity.Task
	res := t.db.Model(TaskModel{}).Find(tt).Where("id = ?", id).Update("open", true)
	if res.Error != nil {
		return res.Error
	}

	return nil
}
