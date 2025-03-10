package main

import (
	"fmt"
	"github/LissaiDev/spl-auth/internal/auth"
	"log"

	"github.com/google/uuid"
)

func main() {
	repo := auth.GetUserRepository()
	user, err := repo.CreateUser(&auth.User{
		Name:     "lissaidev",
		Email:    "lissaidev@gmail.com",
		Password: "Abc123!@",
	})
	fmt.Println(user.ID)
	if err != nil {
		panic(err)
	}
	reUser, reErr := repo.GetUserByID(uuid.New())
	if reErr != nil {
		fmt.Println(reErr)
	}
	log.Printf("Retrieved user: %+v", reUser)
}
