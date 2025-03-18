package auth

import (
	"github/LissaiDev/spl-auth/middleware"

	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.Engine, userController UserController) {
	userGroup := r.Group("/auth")
	userGroup.POST("/", userController.createUser)
	userGroup.GET("/", middleware.Auth(), userController.getUser)
	userGroup.POST("/login", userController.authenticateUser)
}
