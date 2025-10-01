package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AdminIndex 管理页面首页 - 重定向到静态文件
func AdminIndex(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/admin/static/")
}
