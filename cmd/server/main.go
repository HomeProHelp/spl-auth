package main

import (
	"github/LissaiDev/spl-auth/internal/auth"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Initializing
	userRepository := auth.NewUserRepository()
	userService := auth.NewUserService(userRepository)
	userController := auth.NewUserController(userService)

	auth.UserRouter(router, *userController)

	router.Run(":8080")
}
