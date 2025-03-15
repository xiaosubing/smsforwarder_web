package main

import (
	"bg/models"
	"bg/router"
)

func main() {
	// init db
	models.NewDB()

	r := router.App()

	r.Run(":801")
}
