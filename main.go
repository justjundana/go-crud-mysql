package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/justjundana/go-crud-mysql/database/migrations"
	"github.com/justjundana/go-crud-mysql/database/seeders"
	"github.com/justjundana/go-crud-mysql/helper"
	"github.com/justjundana/go-crud-mysql/routes"
)

func main() {
	connectionString := fmt.Sprintf("root:@tcp(localhost:3306)/crud_go?charset=utf8&parseTime=True&loc=Local")
	db, err := helper.InitDB(connectionString)
	if err != nil {
		panic(err)
	}

	flag.Parse()
	args := flag.Args()
	if len(args) >= 1 {
		switch args[0] {
		case "migrate":
			migrations.ExecuteMigrate(db)
			os.Exit(0)
		case "migrate:fresh":
			migrations.DropMigrate(db)
			migrations.ExecuteMigrate(db)
			os.Exit(0)
		case "seed":
			seeders.ExecuteSeeder(db, args[1:]...)
			os.Exit(0)
		}
	}

	e := routes.Router(db)
	e.Logger.Fatal(e.Start(":8000"))
}
