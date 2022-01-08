package main

import (
	"fmt"

	"github.com/justjundana/go-crud-mysql/helper"
	"github.com/justjundana/go-crud-mysql/routes"
)

func main() {
	connectionString := fmt.Sprintf("root:@tcp(localhost:3306)/crud_go?charset=utf8&parseTime=True&loc=Local")
	db, err := helper.InitDB(connectionString)
	if err != nil {
		panic(err)
	}
	e := routes.Router(db)
	e.Logger.Fatal(e.Start(":8000"))
}
