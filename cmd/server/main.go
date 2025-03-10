package main

import (
	"github/LissaiDev/spl-auth/internal/auth"
	"log"
)

func main() {
	repo := auth.GetUserRepository()
	user, err := repo.CreateUser("lissaidev", "lissaidev@gmail.com", "lissaidevpassword")
	if err != nil {
		panic(err)
	}
	reUser, reErr := repo.GetUserByIdentifier(user.Identifier)
	if reErr != nil {
		panic(reErr)
	}
	log.Printf("Retrieved user: %+v", reUser)
}
