package main

import (
	"GoDbConnect/db"
	"GoDbConnect/handler"
	"log"
)

func main() {
	if db.CheckConnection() == 0 {
		log.Fatal("Failed connection with the database")
		return
	}
	handler.Handlers()
}
