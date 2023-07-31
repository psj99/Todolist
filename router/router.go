package router

import (
	"Todolist/api"
	"Todolist/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	ginRouter := gin.Default()
	ginRouter.Use(middleware.Cors())
	store := cookie.NewStore([]byte("something-very-secret"))
	ginRouter.Use(sessions.Sessions("mysession", store))
	v1 := ginRouter.Group("/api/v1")
	{
		v1.GET("ping", func(context *gin.Context) {
			context.JSON(200, "success")
		})

		// 用户操作
		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)

		authed := v1.Group("/") // 需要登陆认证
		authed.Use(middleware.JWT())
		{
			authed.POST("task", api.TaskCreate)
			authed.GET("task", api.TaskList)
			authed.GET("task/:id", api.TaskGet)
			authed.PUT("task", api.UpdateTask)
			authed.POST("task/search", api.SearchTask)
			authed.DELETE("task/:id", api.DeleteTask)
		}

	}
	return ginRouter
}
