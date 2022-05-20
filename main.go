package main

import (
	"github.com/MEVLEPEIZ/Mini-Project/config"
	"github.com/MEVLEPEIZ/Mini-Project/routes"
)

// ---------------------------------------------------
func main() {

	config.InitDB()

	e := routes.NewRoutes()

	e.Logger.Fatal(e.Start(":8080"))

}