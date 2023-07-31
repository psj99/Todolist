package api

import (
	"Todolist/pkg/utils"
	"Todolist/service"
	"Todolist/types"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func TaskCreate(ctx *gin.Context) {
	var req types.CreateTaskReq
	if err := ctx.ShouldBind(&req); err != nil {
		// 参数校验
		utils.ZapLogger.Infoln(err)
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}
	taskSrv := service.GetTaskSrv()
	resp, err := taskSrv.CreateTask(ctx.Request.Context(), &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, resp)

}

func TaskList(ctx *gin.Context) {
	limitstr := ctx.DefaultQuery("limit", "5")
	startstr := ctx.DefaultQuery("start", "1")
	limit, _ := strconv.Atoi(limitstr)
	start, _ := strconv.Atoi(startstr)

	taskSrv := service.GetTaskSrv()
	resp, err := taskSrv.ListTask(ctx.Request.Context(), &types.ListTaskReq{
		Limit: limit,
		Start: start,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func TaskGet(ctx *gin.Context) {
	idstr := ctx.Param("id")
	id, _ := strconv.Atoi(idstr)
	taskSrv := service.GetTaskSrv()
	resp, err := taskSrv.GetTask(ctx.Request.Context(), uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func UpdateTask(ctx *gin.Context) {
	var req types.UpdateTaskReq
	// 参数校验
	if err := ctx.ShouldBind(&req); err == nil {
		utils.ZapLogger.Infoln(err)
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}
	taskSrv := service.GetTaskSrv()
	resp, err := taskSrv.UpdateTask(ctx.Request.Context(), &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func DeleteTask(ctx *gin.Context) {
	idstr := ctx.Param("id")
	id, _ := strconv.Atoi(idstr)
	taskSrv := service.GetTaskSrv()
	resp, err := taskSrv.DeleteTaskById(ctx.Request.Context(), uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func SearchTask(ctx *gin.Context) {
	var req types.SearchTaskReq
	fmt.Printf("%#v : \n", req)
	if err := ctx.ShouldBind(&req); err == nil {
		utils.ZapLogger.Infoln(err)
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}
	taskSrv := service.GetTaskSrv()
	resp, err := taskSrv.SearchTask(ctx.Request.Context(), &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, resp)
}
