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

func (taskSrv *taskSrv) ListTask(ctx context.Context, req *types.ListTaskReq) (resp interface{}, err error) {
	// 从context中获取用户信息
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		utils.ZapLogger.Info(err)
		return
	}

	tasks, total, err := dao.NewTaskDao(ctx).ListTask(req.Start, req.Limit, u.Id)
	if err != nil {
		utils.ZapLogger.Info(err)
		return
	}

	taskRespList := make([]*types.TaskResp, 0)
	for _, v := range tasks {
		taskRespList = append(taskRespList, &types.TaskResp{
			ID:        v.ID,
			Title:     v.Title,
			Content:   v.Content,
			View:      v.View(ctx),
			Status:    v.Status,
			CreatedAt: v.CreatedAt.Unix(),
			StartTime: v.StartTime,
			EndTime:   v.EndTime,
		})
	}

	utils.ZapLogger.Info(taskRespList)

	data := struct {
		Total    int64       `json:"total"`
		TaskList interface{} `json:"list"`
	}{
		Total:    total,
		TaskList: taskRespList,
	}
	return ctl.RespSuccessWithData(data), nil
}

func (taskSrc *taskSrv) GetTask(ctx context.Context, id uint) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		utils.ZapLogger.Info(err)
		return
	}

	task, err := dao.NewTaskDao(ctx).FindTaskByIdAndUserId(u.Id, id)
	if err != nil {
		utils.ZapLogger.Info(err)
		return
	}

	respTask := &types.TaskResp{
		ID:        task.ID,
		Title:     task.Title,
		Content:   task.Content,
		View:      task.View(ctx),
		Status:    task.Status,
		CreatedAt: task.CreatedAt.Unix(),
		StartTime: task.StartTime,
		EndTime:   task.EndTime,
	}

	task.AddView(ctx) // 增加点击数
	return ctl.RespSuccessWithData(respTask), nil

}

func (taskSrc *taskSrv) DeleteTaskById(ctx context.Context, id uint) (resp interface{}, err error) {

	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		utils.ZapLogger.Info(err)
		return
	}
	err = dao.NewTaskDao(ctx).DeleteTaskById(u.Id, id)
	if err != nil {
		utils.ZapLogger.Info(err)
		return
	}

	return ctl.RespSuccess(), nil

}

func (taskSrc *taskSrv) UpdateTask(ctx context.Context, req *types.UpdateTaskReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		utils.ZapLogger.Info(err)
		return
	}
	err = dao.NewTaskDao(ctx).UpdateTask(u.Id, req)
	if err != nil {
		utils.ZapLogger.Info(err)
		return
	}
	return ctl.RespSuccess(), nil
}

func (taskSrc *taskSrv) SearchTask(ctx context.Context, req *types.SearchTaskReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		utils.ZapLogger.Info(err)
		return
	}
	tasks, err := dao.NewTaskDao(ctx).SearchTask(u.Id, req.Info)
	if err != nil {
		utils.ZapLogger.Info(err)
		return
	}
	taskRespList := make([]*types.TaskResp, 0)
	for _, v := range tasks {
		taskRespList = append(taskRespList, &types.TaskResp{
			ID:        v.ID,
			Title:     v.Title,
			Content:   v.Content,
			Status:    v.Status,
			View:      v.View(ctx),
			CreatedAt: v.CreatedAt.Unix(),
			StartTime: v.StartTime,
			EndTime:   v.EndTime,
		})
	}
	return ctl.RespList(taskRespList, int64(len(taskRespList))), nil
}
