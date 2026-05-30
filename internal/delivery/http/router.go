package http

import (
	"net/http"
	docs "task-management-system/docs"
	"task-management-system/internal/delivery/middleware"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(handler Handlers) *gin.Engine {
	r := gin.Default()

	docs.SwaggerInfo.BasePath = "/api"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, ginSwagger.InstanceName("swagger")))

	api := r.Group("/api")
	{
		api.Use(middleware.AuthMiddleware())
		api.GET("/ping", PingHandler)
		UserRouter(api, handler.User)
	}

	return r
}

// PingExample godoc
// @Summary ping
// @Description do the ping
// @Tags PING
// @Accept json
// @Produce json
// @Security     BearerAuth
// @Success 200 {string} string "pong"
// @Router /ping [get]
func PingHandler(g *gin.Context) {
	g.JSON(http.StatusOK, "pong")
}
