package main

import "go-auth-backend/database"

func main() {
	database.Connect("root:password@tcp(localhost:3306)/jwt_demo?parseTime=true")
	database.Migrate()
}
