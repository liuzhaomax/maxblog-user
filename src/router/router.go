package router

import (
	"github.com/gin-gonic/gin"
	"github.com/liuzhaomax/maxblog-user/internal/middleware"
	"github.com/liuzhaomax/maxblog-user/src/api_user/handler"
)

func Register(root *gin.RouterGroup, handler *handler.HandlerUser, mw *middleware.Middleware) {
	root.GET("/login", handler.GetPuk)
	root.POST("/login", handler.PostLogin)
	root.DELETE("/login", handler.DeleteLogin)
	routerUser := root.Group("/users")
	{
		routerUser.GET("/:userID", handler.GetUserByUserID)
	}
}
