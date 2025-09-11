package main

import "example.url/DB/internal/database"

func main() {
	database.PQInitialize(".env")
}
