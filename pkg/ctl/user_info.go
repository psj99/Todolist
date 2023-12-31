package ctl

import (
	"context"
	"errors"
)

type key int

var userKey key

type UserInfo struct {
	Id       uint   `json:"id"`
	UserName string `json:"user_name"`
}

func GetUserInfo(ctx context.Context) (*UserInfo, error) {
	user, ok := FromContext(ctx)
	if !ok {
		return nil, errors.New("获取用户信息错误")
	}
	// ctx.Value("user_id").
	return user, nil
}

func NewContext(ctx context.Context, u *UserInfo) context.Context {

	return context.WithValue(ctx, userKey, u)
}

func FromContext(ctx context.Context) (*UserInfo, bool) {
	u, ok := ctx.Value(userKey).(*UserInfo)
	return u, ok
}
