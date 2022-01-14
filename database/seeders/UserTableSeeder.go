package seeders

import (
	"time"

	"github.com/bxcodec/faker/v3"
)

func (s DatabaseSeeder) UserSeedTableSeeder() {
	for i := 0; i < 100; i++ {
		stmt, _ := s.db.Prepare(`INSERT INTO users(name, email, password, created_at, updated_at) VALUES (?, ? ,?, ?, ?)`)
		_, err := stmt.Exec(faker.Name(), faker.Email(), faker.Password(), time.Now(), time.Now())
		if err != nil {
			panic(err)
		}
	}
}
