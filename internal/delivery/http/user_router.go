package http

import "github.com/gin-gonic/gin"

func UserRouter(rg *gin.RouterGroup, handler *UserHandler) {
	user := rg.Group("/user")
	{
		user.POST("/login", handler.LoginHandler)
		user.POST("/register", handler.RegisterHandler)
	}
}
