package main

import (
	"log"

	"github.com/KenethSandoval/goTwitter/db"
	"github.com/KenethSandoval/goTwitter/handlers"
)

func main() {
	if db.ConnectionCheck() == 0 {
		log.Fatal("Sin conexion a la DB")
		return
	}

	handlers.Handlers()
}
