package seeders

import (
	"database/sql"
	"log"
	"reflect"
)

type DatabaseSeeder struct {
	db *sql.DB
}

func ExecuteSeeder(db *sql.DB, seedMethodNames ...string) {
	s := DatabaseSeeder{db}

	seedType := reflect.TypeOf(s)

	if len(seedMethodNames) == 0 {
		log.Println("Running all seeder...")
		for i := 0; i < seedType.NumMethod(); i++ {
			method := seedType.Method(i)
			Seeder(s, method.Name)
		}
	}

	for _, item := range seedMethodNames {
		Seeder(s, item)
	}
}

func Seeder(s DatabaseSeeder, seedMethodName string) {
	m := reflect.ValueOf(s).MethodByName(seedMethodName)
	if !m.IsValid() {
		log.Fatal("No method called ", seedMethodName)
	}

	log.Println("Seeding", seedMethodName, "...")

	m.Call(nil)
	log.Println("Seeded", seedMethodName, "success")
}
