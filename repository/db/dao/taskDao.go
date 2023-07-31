package dao

import (
	"Todolist/repository/db/model"
	"Todolist/types"
	"context"

	"gorm.io/gorm"
)

type taskDao struct {
	*gorm.DB
}

func NewTaskDao(ctx context.Context) *taskDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &taskDao{NewDBClient(ctx)}
}

// CreateTask 创建Task
func (dao *taskDao) CreateTask(task *model.Task) error {
	return dao.Model(&model.Task{}).Create(&task).Error
}

// ListTask List Task的情况
func (dao *taskDao) ListTask(start, limit int, userId uint) (r []*model.Task, total int64, err error) {
	err = dao.Model(&model.Task{}).Preload("User").Where("uid = ?", userId).
		Count(&total).
		Limit(limit).Offset((start - 1) * limit).
		Find(&r).Error

	return
}

// FindTaskByIdAndUserId 通过id和user_id找到task
func (dao *taskDao) FindTaskByIdAndUserId(uId, id uint) (r *model.Task, err error) {
	err = dao.Model(&model.Task{}).Where("id = ? AND uid = ?", id, uId).First(&r).Error

	return
}

// UpdateTask 修改
func (s *taskDao) UpdateTask(uId uint, req *types.UpdateTaskReq) error {
	t := new(model.Task)
	err := s.Model(&model.Task{}).Where("id = ? AND uid=?", req.ID, uId).First(&t).Error
	if err != nil {
		return err
	}

	if req.Status != 0 {
		t.Status = req.Status
	}

	if req.Title != "" {
		t.Title = req.Title
	}

	if req.Content != "" {
		t.Content = req.Content
	}

	return s.Save(t).Error
}

// SearchTask 搜索Task
func (dao *taskDao) SearchTask(uId uint, info string) (tasks []*model.Task, err error) {

	err = dao.Where("uid=?", uId).Preload("User").First(&tasks).Error
	if err != nil {
		return
	}

	err = dao.Model(&model.Task{}).Where("title LIKE ? OR content LIKE ?",
		"%"+info+"%", "%"+info+"%").Find(&tasks).Error

	return
}

// DeleteTaskById 通过id删除
func (dao *taskDao) DeleteTaskById(uId, tId uint) error {
	r, err := dao.FindTaskByIdAndUserId(uId, tId)
	if err != nil {
		return err
	}
	return dao.Delete(&r).Error
}
