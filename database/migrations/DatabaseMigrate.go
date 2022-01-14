package migrations

import (
	"database/sql"
	"path"
	"runtime"
)

func ExecuteMigrate(db *sql.DB) {
	ImportUserDB(db)
	ImportBookDB(db)
	ImportProductDB(db)
}

func DropMigrate(db *sql.DB) {
	DropProductDB(db)
	DropBookDB(db)
	DropUserDB(db)
}

func GetSourcePath() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}
