package main

import (
	"log"
	"os"

	"github.com/vivaldy22/enigma_bank/config"
)

func main() {
	args := os.Args

	if len(args) != 3 {
		log.Println("don't forget to run with arguments, 1st is host, 2nd is port")
		return
	}

	db, _ := config.InitDB()
	router := config.CreateRouter()
	config.InitRouters(db, router)
	config.RunServer(router, args[1], args[2])
}
