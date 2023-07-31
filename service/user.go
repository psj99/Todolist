package service

import (
	"Todolist/pkg/ctl"
	"Todolist/pkg/utils"
	"Todolist/repository/db/dao"
	"Todolist/repository/db/model"
	"Todolist/types"
	"context"
	"errors"
	"sync"

	"gorm.io/gorm"
)

var userSrvIns *userSrv
var userSrvOnce sync.Once

type userSrv struct {
}

func GetUserService() *userSrv {
	userSrvOnce.Do(func() {
		userSrvIns = &userSrv{}
	})
	return userSrvIns
}

func (s *userSrv) Register(ctx context.Context, req *types.UserRegisterReq) (resp interface{}, err error) {
	userDao := dao.NewUserDao(ctx)
	_, err = userDao.FindUserByUserName(req.UserName)
	switch err { // TODO 优化一下
	case gorm.ErrRecordNotFound:
		u := &model.User{
			UserName: req.UserName,
		}
		// 密码加密存储
		if err = u.SetPassword(req.Password); err != nil {
			utils.ZapLogger.Info(err)
			return
		}

		if err = userDao.CreateUser(u); err != nil {
			utils.ZapLogger.Info(err)
			return
		}

		return ctl.RespSuccess(), nil
	case nil:
		err = errors.New("用户已存在")
		return
	default:
		return
	}
}

// Login 用户登陆函数
func (s *userSrv) Login(ctx context.Context, req *types.UserLoginReq) (resp interface{}, err error) {
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.FindUserByUserName(req.UserName)
	if err == gorm.ErrRecordNotFound {
		err = errors.New("用户不存在")
		return
	}

	if !user.CheckPassword(req.Password) {
		err = errors.New("账号/密码错误")
		utils.ZapLogger.Info(err)
		return
	}

	token, err := utils.GenerateToken(user.ID, user.UserName)
	if err != nil {
		utils.ZapLogger.Infoln(err)
		return
	}

	userLoginResp := &types.TokenData{
		User: &types.UserInfoResp{
			ID:       user.ID,
			UserName: user.UserName,
			CreateAt: user.CreatedAt.Unix(),
		},
		AccessToken: token,
	}

	return ctl.RespSuccessWithData(userLoginResp), nil
}
