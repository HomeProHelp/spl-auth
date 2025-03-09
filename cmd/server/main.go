package main

import "github/LissaiDev/spl-auth/internal/auth"

func main() {
	auth.GetUserRepository().CreateUser("lissaidev", "lissaidev@gmail.com", "lissaidevpassword")
}
