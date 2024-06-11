package main

import (
	"log"
	api "pkg/FrontendConnection"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	log.Print("[SERVER] Server running on http://127.0.0.1:8080")
	api.RunServer()
}
