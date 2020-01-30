package server

import (
	"os"
	"singo/api"
	"singo/middleware"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	//跨域
	r.Use(middleware.Cors())
	//验证登录
	r.Use(middleware.CurrentUser())

	// 路由
	v1 := r.Group("/api/v1")
	{
		//restful
		//创建视频，投稿
		v1.POST("videos", api.CreateVideo)

		//获得视频详情
		v1.GET("videos",api.ListVideo)

		//视频列表
		v1.GET("video/:id",api.ShowVideo)

		//更新视频
		v1.PUT("video/:id",api.UpdateVideo)

		//删除视频
		v1.DELETE("video/:id",api.DeleteVideo)

		//检测生存状态
		v1.POST("ping", api.Ping)

		// 用户注册
		v1.POST("user/register", api.UserRegister)

		// 用户登录
		v1.POST("user/login", api.UserLogin)

		// 需要登录保护的
		auth := v1.Group("")
		auth.Use(middleware.AuthRequired())
		{
			// User Routing
			auth.GET("user/me", api.UserMe)
			auth.DELETE("user/logout", api.UserLogout)
		}

	}
	return r
}
