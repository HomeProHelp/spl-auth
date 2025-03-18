package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	usr, error := ctrl.srv.CreateUser(&user)
	if error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": error.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, usr)
	return
}
