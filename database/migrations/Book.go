package migrations

import (
	"database/sql"
	"io/ioutil"
)

func ImportBookDB(db *sql.DB) {
	q, err := ioutil.ReadFile(GetSourcePath() + "/sql/books.sql")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(string(q))
	if err != nil {
		panic(err)
	}
}

func DropBookDB(db *sql.DB) {
	q := "DROP TABLE books;"

	_, err := db.Exec(string(q))
	if err != nil {
		panic(err)
	}
}
