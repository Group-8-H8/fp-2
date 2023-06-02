package main

import (
	"fmt"
	"github.com/Group-8-H8/fp-2.git/config"
	"github.com/Group-8-H8/fp-2.git/db"
	"github.com/Group-8-H8/fp-2.git/router"
)

func main() {
	r := router.StartApp()
	err := database.StartDB()
	if err != nil {
		fmt.Println("Error starting database: ", err)
		return
	}
	r.Run(config.SERVER_PORT)
}