package middleware

import (
	"Todolist/pkg/ctl"
	"Todolist/pkg/e"
	"Todolist/pkg/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		code := e.SUCCESS
		token := ctx.GetHeader("Authorization")
		if token == "" {
			code = http.StatusNotFound
			ctx.JSON(e.InvalidParams, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
				"data":   "缺少token",
			})
			ctx.Abort()
			return
		}

		claims, err := utils.ParseToken(token)
		if err != nil {
			code = e.ErrorAuthCheckTokenFail
		} else if time.Now().Unix() > claims.ExpiresAt.Unix() {
			code = e.ErrorAuthCheckTokenTimeout
		}

		if code != e.SUCCESS {
			ctx.JSON(e.InvalidParams, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
				"data":   "可能是身份过期了，请重新登录",
			})
			ctx.Abort()
			return
		}

		// ctx.Set("user_id", claims.ID)
		ctx.Request = ctx.Request.WithContext(ctl.NewContext(ctx.Request.Context(),
			&ctl.UserInfo{Id: claims.Id, UserName: claims.UserName}))
		ctx.Next()
	}
}
