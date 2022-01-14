package migrations

import (
	"database/sql"
	"io/ioutil"
)

func ImportProductDB(db *sql.DB) {
	q, err := ioutil.ReadFile(GetSourcePath() + "/sql/products.sql")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(string(q))
	if err != nil {
		panic(err)
	}
}

func DropProductDB(db *sql.DB) {
	q := "DROP TABLE products;"

	_, err := db.Exec(string(q))
	if err != nil {
		panic(err)
	}
}
