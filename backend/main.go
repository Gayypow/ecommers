package main

import (
	"github.com/MilliGoshant/merci/backend/routes"

	"github.com/MilliGoshant/merci/backend/db"
)

func init() {
	db.Connect()
}

func main() {
	r := routes.SetupRoutes()
	r.Run()
}
