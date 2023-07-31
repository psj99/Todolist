package types

type UserRegisterReq struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=3,max=15" example:"FanOne"`
	Password string `form:"password" json:"password" binding:"required,min=5,max=16" example:"FanOne666"`
}
type UserLoginReq struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=3,max=15" example:"FanOne"`
	Password string `form:"password" json:"password" binding:"required,min=5,max=16" example:"FanOne666"`
}

type UserInfoResp struct {
	ID       uint   `json:"id" form:"id" example:"1"`                    // 用户ID
	UserName string `json:"user_name" form:"user_name" example:"FanOne"` // 用户名
	CreateAt int64  `json:"create_at" form:"create_at"`                  // 创建
}

// TokenData 带有token的Data结构
type TokenData struct {
	User         interface{} `json:"user"`
	AccessToken  string      `json:"access_token"`
	RefreshToken string      `json:"refresh_token"`
}
