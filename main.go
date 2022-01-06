package main

import (
	"crud-database/helper"
	"crud-database/routes"
)

func main() {
	helper.InitDB()
	e := routes.New()
	e.Logger.Fatal(e.Start(":8000"))
}
