package handlers

import (
	"os"

	"github.com/gin-gonic/gin"
)

// BasicAuth 中间件，从环境变量获取用户名和密码
func BasicAuth() gin.HandlerFunc {
	username := os.Getenv("ADMIN_USERNAME")
	password := os.Getenv("ADMIN_PASSWORD")

	// 如果未设置环境变量，使用默认值
	if username == "" {
		username = "admin"
	}
	if password == "" {
		password = "password"
	}

	return gin.BasicAuth(gin.Accounts{
		username: password,
	})
}
