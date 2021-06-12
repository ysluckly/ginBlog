package proxy

import (
	"net/http"

	"ginBlog/api/v1"
	"ginBlog/utils"

	"github.com/gin-gonic/gin"
)

// InitRouter 初始化路由
func InitRouter(){
	gin.SetMode(utils.AppMode)

	r := gin.Default()
	router1 := r.Group("api/blog/app/v1")
	{
		// 测试专用模块
		router1.GET("/test", func(c *gin.Context) {

			c.JSON(http.StatusOK, gin.H{
				"status": "success",
			})
		})

		// 用户信息模块
		router1.POST("user/add", v1.AddUser)
		router1.GET("user/:id", v1.GetUserInfo)
		router1.GET("users", v1.GetUsers)

		/*// 文章分类信息模块
		router1.GET("category", v1.GetCate)
		router1.GET("category/:id", v1.GetCateInfo)

		// 文章模块
		router1.GET("article", v1.GetArt)
		router1.GET("article/list/:id", v1.GetCateArt)
		router1.GET("article/info/:id", v1.GetArtInfo)

		// 登录控制模块
		router1.POST("login", v1.Login)
		router1.POST("login/front", v1.LoginFront)

		// 获取个人设置信息
		router1.GET("profile/:id", v1.GetProfile)

		// 评论模块
		router1.POST("comment/add", v1.AddComment)
		router1.GET("comment/info/:id", v1.GetComment)
		router1.GET("comment/front/:id", v1.GetCommentListFront)
		router1.GET("comment/count/:id", v1.GetCommentCount)*/
	}

	router2 := r.Group("api/blog/admin/v1")
	{
		//*********************** 测试模块 *************************/
		router2.GET("/test", func(c *gin.Context) {

			c.JSON(http.StatusOK, gin.H{
				"status": "success",
			})
		})
		// 用户模块
		router2.GET("admin/users", v1.GetUsers)
		router2.PUT("user/:id", v1.EditUser)
		router2.DELETE("user/:id", v1.DeleteUser)

	}
	_  = r.Run(utils.HttpPort)
}
