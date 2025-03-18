package auth

import "github.com/gin-gonic/gin"

func UserRouter(r *gin.Engine, userController UserController) {
	userGroup := r.Group("/users")
	userGroup.POST("/", userController.createUser)
	userGroup.GET("/:id", userController.getUser)
}
