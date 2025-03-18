package auth

import (
	"github/LissaiDev/spl-auth/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserController struct {
	srv *UserService
}

func NewUserController(srv *UserService) *UserController {
	return &UserController{
		srv: srv,
	}
}

func (ctrl *UserController) createUser(ctx *gin.Context) {
	var user User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, &utils.Response{Code: utils.AuthenticationCodes["invalid_data"], Data: map[string]string{}})
		return
	}
	usr, code := ctrl.srv.CreateUser(&user)
	if !utils.IsSuccess(code) {
		ctx.JSON(http.StatusInternalServerError, &utils.Response{Code: code, Data: map[string]string{}})
		return
	}
	ctx.JSON(http.StatusCreated, &utils.Response{Code: utils.AuthenticationCodes["success"], Data: utils.UserResponse{
		ID:        usr.ID.String(),
		Name:      usr.Name,
		Email:     usr.Email,
		CreatedAt: usr.CreatedAt,
		UpdatedAt: usr.UpdatedAt,
	}})
	return
}

func (ctrl *UserController) getUser(ctx *gin.Context) {
	id, exists := ctx.Get("ID")
	if !exists {
		ctx.JSON(http.StatusBadRequest, &utils.Response{Code: utils.AuthenticationCodes["invalid_data"], Data: map[string]string{}})
		return
	}

	parsedUUID, err := uuid.Parse(id.(string))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &utils.Response{Code: utils.AuthenticationCodes["invalid_data"], Data: map[string]string{}})
		return
	}

	user, code := ctrl.srv.GetUser(parsedUUID)
	if !utils.IsSuccess(code) {

		if code == utils.AuthenticationCodes["user_not_found"] {
			ctx.JSON(http.StatusNotFound, &utils.Response{Code: utils.AuthenticationCodes["user_not_found"], Data: map[string]string{}})
			return
		}

		ctx.JSON(http.StatusInternalServerError, &utils.Response{Code: code, Data: map[string]string{}})
		return
	}
	ctx.JSON(http.StatusOK, &utils.Response{Code: utils.AuthenticationCodes["success"], Data: utils.UserResponse{
		ID:        user.ID.String(),
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}})
	return
}

func (ctrl *UserController) authenticateUser(ctx *gin.Context) {
	var user utils.UserRequest
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, &utils.Response{Code: utils.AuthenticationCodes["invalid_data"], Data: map[string]string{}})
		return
	}

	token, code := ctrl.srv.AuthenticateUser(user.Email, user.Password)

	if !utils.IsSuccess(code) {
		ctx.JSON(http.StatusBadRequest, &utils.Response{Code: code, Data: map[string]string{}})
		return
	}

	ctx.JSON(http.StatusOK, &utils.Response{Code: utils.AuthenticationCodes["success"], Data: map[string]string{
		"token": *token,
	}})
	return
}
