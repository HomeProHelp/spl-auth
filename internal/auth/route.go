package auth

import "github.com/gin-gonic/gin"

func UserRouter(r *gin.Engine, userController UserController) {
	userGroup := r.Group("/auth")
	userGroup.POST("/", userController.createUser)
	userGroup.GET("/:id", userController.getUser)
	userGroup.POST("/login", userController.authenticateUser)
}
