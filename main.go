package main

import (
	"jwt-go-system/controllers"
	"jwt-go-system/initializers"
	"jwt-go-system/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	// 加载env文件里的变量
	initializers.LoadEnvVariables()
	// 链接数据库
	initializers.ConnectToDb()
	// 建表
	initializers.SyncDatabase()
}
func main() {
	r := gin.Default()
	r.Use(middleware.Core())
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	r.Run(":3000")
}
