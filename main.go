// main.go

package main

import (
	"fmt"
	"log"
	"tonx/pkg/config"
	"tonx/pkg/db"
	"tonx/routers"
)

func main() {
	log.Println(" Init DB ")
	db.InitDB()

	log.Println(" Start Server ")

	router := routers.InitRouter()
	port := config.Config.Port
	router.Run(fmt.Sprintf(":%d", port))
}
