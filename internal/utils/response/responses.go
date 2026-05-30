package response

import (
	"task-management-system/internal/utils/dto"

	"github.com/gin-gonic/gin"
)

func Succes[T any](g *gin.Context, statusCode int, data T, meta *dto.Meta) {
	g.JSON(statusCode, dto.Response[T]{
		Success: true,
		Data:    data,
		Meta:    meta,
	})
}

func Failed(g *gin.Context, statusCode int, errCode dto.ErrorCode, errMessage string) {
	g.JSON(statusCode, dto.Response[any]{
		Success: false,
		Error: &dto.ErrorInfo{
			Code:    errCode,
			Message: errMessage,
		},
	})
}

func AbortFailed(g *gin.Context, statusCode int, errCode dto.ErrorCode, errMessage string) {
	g.AbortWithStatusJSON(statusCode, dto.Response[any]{
		Success: false,
		Error: &dto.ErrorInfo{
			Code:    errCode,
			Message: errMessage,
		},
	})
}
