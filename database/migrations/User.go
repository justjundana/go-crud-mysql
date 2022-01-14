package migrations

import (
	"database/sql"
	"io/ioutil"
)

func ImportUserDB(db *sql.DB) {
	q, err := ioutil.ReadFile(GetSourcePath() + "/sql/users.sql")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(string(q))
	if err != nil {
		panic(err)
	}
}

func DropUserDB(db *sql.DB) {
	q := "DROP TABLE users;"

	_, err := db.Exec(string(q))
	if err != nil {
		panic(err)
	}
}
