package http

import (
	"net/http"
	"task-management-system/internal/domain"
	"task-management-system/internal/utils/dto"
	"task-management-system/internal/utils/response"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	uc domain.UserUsecase
}

func NewUserHandler(uc domain.UserUsecase) *UserHandler {
	return &UserHandler{uc: uc}
}

// LoginHandler godoc
// @Summary      User Login
// @Description  Autentikasi user menggunakan email dan password untuk mendapatkan JWT token
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        request  body      dto.LoginRequest  true  "Payload Login"
// @Success      200      {object}  dto.Response[dto.LoginResponse]
// @Failure      400      {object}  dto.Response[any]
// @Router       /user/login [post]
func (h *UserHandler) LoginHandler(g *gin.Context) {
	var reqLogin dto.LoginRequest
	if err := g.ShouldBindJSON(&reqLogin); err != nil {
		response.Failed(g, http.StatusBadRequest, dto.ErrInvalidInput, err.Error())
		return
	}

	loginResp, err := h.uc.Login(g.Request.Context(), reqLogin)
	if err != nil {
		if custErr, ok := dto.AsType[dto.ErrorCode](err); ok {
			response.Failed(g, http.StatusBadRequest, custErr, err.Error())
			return
		}
		response.Failed(g, http.StatusInternalServerError, dto.ErrInternalServer, err.Error())
		return
	}

	response.Succes(g, http.StatusOK, loginResp, nil)
}

// RegisterHandler godoc
// @Summary      User Register
// @Description  Autentikasi user menggunakan email dan password untuk mendapatkan JWT token
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        request  body      dto.RegisterRequest  true  "Payload Register"
// @Success      200      {object}  dto.Response[dto.RegisterResponse]
// @Failure      400      {object}  dto.Response[any]
// @Router       /user/register [post]
func (h *UserHandler) RegisterHandler(g *gin.Context) {
	var reqRegister dto.RegisterRequest
	if err := g.ShouldBindBodyWithJSON(&reqRegister); err != nil {
		response.Failed(g, http.StatusBadRequest, dto.ErrInvalidInput, err.Error())
		return
	}

	registerResp, err := h.uc.Register(g.Request.Context(), reqRegister)
	if err != nil {
		if custErr, ok := dto.AsType[dto.ErrorCode](err); ok {
			response.Failed(g, http.StatusBadRequest, custErr, err.Error())
		}
		response.Failed(g, http.StatusInternalServerError, dto.ErrInternalServer, err.Error())
	}

	response.Succes(g, http.StatusOK, registerResp, nil)
}
