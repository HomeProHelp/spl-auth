package main

import "github/LissaiDev/spl-auth/db"

func main() {
	db := db.GetDatabaseInstance()
	print(db.AllowGlobalUpdate)
}
