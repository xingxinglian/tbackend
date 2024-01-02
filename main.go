// main.go

package main

import (
	"tonx/pkg/db"
	"tonx/routers"
)

func main() {
	db.InitDB()

	router := routers.InitRouter()
	router.Run(":80")
}
