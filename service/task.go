package service

import (
	"Todolist/pkg/ctl"
	"Todolist/pkg/utils"
	"Todolist/repository/db/dao"
	"Todolist/repository/db/model"
	"Todolist/types"
	"context"
	"sync"
	"time"
)

var taskSrvIns *taskSrv
var taskSrvOnce sync.Once

type taskSrv struct {
}

func GetTaskSrv() *taskSrv {
	taskSrvOnce.Do(func() {
		taskSrvIns = &taskSrv{}
	})
	return taskSrvIns
}

func (taskSrv *taskSrv) CreateTask(ctx context.Context, req *types.CreateTaskReq) (resp interface{}, err error) {
	// 从context中获取用户信息
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		utils.ZapLogger.Info(err)
		return
	}
	user, err := dao.NewUserDao(ctx).FindUserByUserId(u.Id)
	if err != nil {
		utils.ZapLogger.Info(err)
		return
	}

	task := &model.Task{
		Uid:       user.ID,
		Title:     req.Title,
		Content:   req.Content,
		Status:    0,
		StartTime: time.Now().Unix(),
	}
	err = dao.NewTaskDao(ctx).CreateTask(task)
	if err != nil {
		utils.ZapLogger.Info(err)
		return
	}
	return ctl.RespSuccess(), nil

}
