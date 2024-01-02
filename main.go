// main.go

package main

import (
	"log"
	"tonx/pkg/db"
	"tonx/routers"
)

func main() {
	log.Println(" Init DB ")
	db.InitDB()

	log.Println(" Start Server ")

	router := routers.InitRouter()
	router.Run(":80")
}
