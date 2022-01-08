package main

import (
	"github.com/justjundana/go-crud-mysql/helper"
	"github.com/justjundana/go-crud-mysql/routes"
)

func main() {
	helper.InitDB()
	e := routes.New()
	e.Logger.Fatal(e.Start(":8000"))
}
