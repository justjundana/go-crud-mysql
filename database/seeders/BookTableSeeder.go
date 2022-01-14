package seeders

import (
	"time"

	"github.com/bxcodec/faker/v3"
)

func (s DatabaseSeeder) BookSeedTableSeeder() {
	for i := 0; i < 100; i++ {
		stmt, _ := s.db.Prepare(`INSERT INTO books(code, title, description, author, publisher, created_at, updated_at) VALUES (?, ? ,?, ?, ?, ?, ?)`)
		_, err := stmt.Exec(faker.CCNumber(), faker.Sentence(), faker.Paragraph(), faker.Name(), faker.Word(), time.Now(), time.Now())
		if err != nil {
			panic(err)
		}
	}
}
