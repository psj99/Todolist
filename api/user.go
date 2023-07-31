package api

import (
	"Todolist/pkg/utils"
	"Todolist/service"
	"Todolist/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRegister(ctx *gin.Context) {
	var req types.UserRegisterReq
	if err := ctx.ShouldBind(&req); err != nil {
		// 参数校验
		utils.ZapLogger.Infoln(err)
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}
	userService := service.GetUserService()
	resp, err := userService.Register(ctx.Request.Context(), &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
	}
	ctx.JSON(http.StatusOK, resp)

}

func UserLogin(ctx *gin.Context) {
	var req types.UserLoginReq

	if err := ctx.ShouldBind(&req); err != nil {
		// 参数校验
		utils.ZapLogger.Infoln(err)
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	userService := service.GetUserService()
	resp, err := userService.Login(ctx.Request.Context(), &req)
	if err != nil {
		ctx.JSON(http.StatusForbidden, ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, resp)

}
