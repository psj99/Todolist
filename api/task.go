package api

import (
	"Todolist/pkg/utils"
	"Todolist/service"
	"Todolist/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TaskCreate(ctx *gin.Context) {

	var req types.CreateTaskReq
	if err := ctx.ShouldBind(&req); err != nil {
		// 参数校验
		utils.ZapLogger.Infoln(err)
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
	taskSrv := service.GetTaskSrv()
	resp, err := taskSrv.CreateTask(ctx.Request.Context(), &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, resp)

}
